package execx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type envMode int

type pipeMode int

const (
	envInherit envMode = iota
	envOnly
	envAppend
)

const (
	pipeStrict pipeMode = iota
	pipeBestEffort
)

// Command constructs a new command without executing it.
func Command(name string, args ...string) *Cmd {
	cmd := &Cmd{
		name:     name,
		args:     append([]string{}, args...),
		envMode:  envInherit,
		pipeMode: pipeStrict,
	}
	cmd.root = cmd
	return cmd
}

// Cmd represents a single command invocation or a pipeline stage.
type Cmd struct {
	name string
	args []string

	env     map[string]string
	envMode envMode
	ctx     context.Context
	cancel  context.CancelFunc
	dir     string

	stdin io.Reader

	onStdout func(string)
	onStderr func(string)

	next     *Cmd
	root     *Cmd
	pipeMode pipeMode
}

// Arg appends arguments to the command.
func (c *Cmd) Arg(values ...any) *Cmd {
	for _, value := range values {
		switch v := value.(type) {
		case string:
			c.args = append(c.args, v)
		case []string:
			c.args = append(c.args, v...)
		case map[string]string:
			keys := make([]string, 0, len(v))
			for key := range v {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			for _, key := range keys {
				c.args = append(c.args, key, v[key])
			}
		default:
			c.args = append(c.args, fmt.Sprint(v))
		}
	}
	return c
}

// Env adds environment variables to the command.
func (c *Cmd) Env(values ...any) *Cmd {
	if c.env == nil {
		c.env = map[string]string{}
	}
	for _, value := range values {
		switch v := value.(type) {
		case string:
			key, val, _ := strings.Cut(v, "=")
			c.env[key] = val
		case []string:
			for _, entry := range v {
				key, val, _ := strings.Cut(entry, "=")
				c.env[key] = val
			}
		case map[string]string:
			for key, val := range v {
				c.env[key] = val
			}
		default:
			key, val, _ := strings.Cut(fmt.Sprint(v), "=")
			c.env[key] = val
		}
	}
	return c
}

// EnvInherit restores default environment inheritance.
func (c *Cmd) EnvInherit() *Cmd {
	c.envMode = envInherit
	return c
}

// EnvOnly ignores the parent environment.
func (c *Cmd) EnvOnly(values map[string]string) *Cmd {
	c.envMode = envOnly
	c.env = map[string]string{}
	for key, val := range values {
		c.env[key] = val
	}
	return c
}

// EnvAppend merges variables into the inherited environment.
func (c *Cmd) EnvAppend(values map[string]string) *Cmd {
	c.envMode = envAppend
	if c.env == nil {
		c.env = map[string]string{}
	}
	for key, val := range values {
		c.env[key] = val
	}
	return c
}

// Dir sets the working directory.
func (c *Cmd) Dir(path string) *Cmd {
	c.dir = path
	return c
}

// WithContext binds the command to a context.
func (c *Cmd) WithContext(ctx context.Context) *Cmd {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	c.ctx = ctx
	return c
}

// WithTimeout binds the command to a timeout.
func (c *Cmd) WithTimeout(d time.Duration) *Cmd {
	parent := c.ctx
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
		parent = nil
	}
	if parent == nil || parent.Err() != nil {
		parent = context.Background()
	}
	c.ctx, c.cancel = context.WithTimeout(parent, d)
	return c
}

// WithDeadline binds the command to a deadline.
func (c *Cmd) WithDeadline(t time.Time) *Cmd {
	parent := c.ctx
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
		parent = nil
	}
	if parent == nil || parent.Err() != nil {
		parent = context.Background()
	}
	c.ctx, c.cancel = context.WithDeadline(parent, t)
	return c
}

// StdinString sets stdin from a string.
func (c *Cmd) StdinString(input string) *Cmd {
	c.stdin = strings.NewReader(input)
	return c
}

// StdinBytes sets stdin from bytes.
func (c *Cmd) StdinBytes(input []byte) *Cmd {
	c.stdin = bytes.NewReader(input)
	return c
}

// StdinReader sets stdin from an io.Reader.
func (c *Cmd) StdinReader(reader io.Reader) *Cmd {
	c.stdin = reader
	return c
}

// StdinFile sets stdin from a file.
func (c *Cmd) StdinFile(file *os.File) *Cmd {
	c.stdin = file
	return c
}

// OnStdout registers a line callback for stdout.
func (c *Cmd) OnStdout(fn func(string)) *Cmd {
	c.onStdout = fn
	return c
}

// OnStderr registers a line callback for stderr.
func (c *Cmd) OnStderr(fn func(string)) *Cmd {
	c.onStderr = fn
	return c
}

// Pipe appends a new command to the pipeline.
func (c *Cmd) Pipe(name string, args ...string) *Cmd {
	root := c.rootCmd()
	next := &Cmd{
		name:     name,
		args:     append([]string{}, args...),
		envMode:  envInherit,
		pipeMode: root.pipeMode,
		root:     root,
	}
	last := root
	for last.next != nil {
		last = last.next
	}
	last.next = next
	return next
}

// PipeStrict sets strict pipeline semantics.
func (c *Cmd) PipeStrict() *Cmd {
	c.rootCmd().pipeMode = pipeStrict
	return c
}

// PipeBestEffort sets best-effort pipeline semantics.
func (c *Cmd) PipeBestEffort() *Cmd {
	c.rootCmd().pipeMode = pipeBestEffort
	return c
}

// Args returns the argv slice used for execution.
func (c *Cmd) Args() []string {
	args := make([]string, 0, len(c.args)+1)
	args = append(args, c.name)
	args = append(args, c.args...)
	return args
}

// EnvList returns the environment list for execution.
func (c *Cmd) EnvList() []string {
	return buildEnv(c.envMode, c.env)
}

// String returns a human-readable representation of the command.
func (c *Cmd) String() string {
	parts := make([]string, 0, len(c.args)+1)
	parts = append(parts, c.name)
	for _, arg := range c.args {
		if strings.ContainsAny(arg, " \t\n\r") {
			parts = append(parts, strconv.Quote(arg))
			continue
		}
		parts = append(parts, arg)
	}
	return strings.Join(parts, " ")
}

// ShellEscaped returns a shell-escaped string for logging only.
func (c *Cmd) ShellEscaped() string {
	parts := make([]string, 0, len(c.args)+1)
	parts = append(parts, shellEscape(c.name))
	for _, arg := range c.args {
		parts = append(parts, shellEscape(arg))
	}
	return strings.Join(parts, " ")
}

// Run executes the command and returns the result.
func (c *Cmd) Run() Result {
	result, _ := c.runInternal(false)
	return result
}

// Output executes the command and returns stdout.
func (c *Cmd) Output() (string, error) {
	result := c.Run()
	return result.Stdout, result.Err
}

// OutputBytes executes the command and returns stdout bytes.
func (c *Cmd) OutputBytes() ([]byte, error) {
	result := c.Run()
	return []byte(result.Stdout), result.Err
}

// OutputTrimmed executes the command and returns trimmed stdout.
func (c *Cmd) OutputTrimmed() (string, error) {
	result := c.Run()
	return strings.TrimSpace(result.Stdout), result.Err
}

// CombinedOutput executes the command and returns stdout+stderr.
func (c *Cmd) CombinedOutput() (string, error) {
	result, combined := c.runInternal(true)
	return combined, result.Err
}

// Start executes the command asynchronously.
func (c *Cmd) Start() *Process {
	proc := &Process{resultCh: make(chan Result, 1)}
	go func() {
		res := c.Run()
		proc.resultCh <- res
		close(proc.resultCh)
	}()
	return proc
}

func (c *Cmd) ctxOrBackground() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

func (c *Cmd) rootCmd() *Cmd {
	if c.root != nil {
		return c.root
	}
	return c
}

type stage struct {
	cmd         *exec.Cmd
	def         *Cmd
	stdoutBuf   bytes.Buffer
	stderrBuf   bytes.Buffer
	combinedBuf bytes.Buffer
	startErr    error
	waitErr     error
	startTime   time.Time
	pipeWriter  *io.PipeWriter
}

func (c *Cmd) runInternal(withCombined bool) (Result, string) {
	stages := c.pipelineStages()
	for _, stage := range stages {
		stage.startTime = time.Now()
		stage.cmd = stage.def.execCmd()
		stdoutWriter := stage.def.stdoutWriter(&stage.stdoutBuf, withCombined, &stage.combinedBuf)
		stderrWriter := stage.def.stderrWriter(&stage.stderrBuf, withCombined, &stage.combinedBuf)
		stage.cmd.Stdout = stdoutWriter
		stage.cmd.Stderr = stderrWriter
	}

	for i := range stages {
		if i == 0 {
			stages[i].cmd.Stdin = stages[i].def.stdin
			continue
		}
		reader, writer := io.Pipe()
		stages[i-1].pipeWriter = writer
		stages[i].cmd.Stdin = reader
		stages[i-1].cmd.Stdout = io.MultiWriter(stages[i-1].cmd.Stdout, writer)
	}

	for i, stage := range stages {
		stage.startErr = stage.cmd.Start()
		if stage.startErr != nil {
			for j := i + 1; j < len(stages); j++ {
				stages[j].startErr = stage.startErr
			}
			break
		}
	}

	for i := range stages {
		if stages[i].startErr != nil {
			if stages[i].pipeWriter != nil {
				_ = stages[i].pipeWriter.Close()
			}
			continue
		}
		stages[i].waitErr = stages[i].cmd.Wait()
		if stages[i].pipeWriter != nil {
			_ = stages[i].pipeWriter.Close()
		}
	}

	results := make([]Result, 0, len(stages))
	for _, stage := range stages {
		results = append(results, stage.result())
	}

	primaryIndex := len(results) - 1
	if c.rootCmd().pipeMode == pipeStrict {
		for i, res := range results {
			if res.ExitCode != 0 || res.Err != nil {
				primaryIndex = i
				break
			}
		}
	}

	primary := results[primaryIndex]
	combined := ""
	if withCombined {
		combined = stages[primaryIndex].combinedBuf.String()
	}
	return primary, combined
}

func (s *stage) result() Result {
	res := Result{
		Stdout:   s.stdoutBuf.String(),
		Stderr:   s.stderrBuf.String(),
		ExitCode: -1,
		Duration: time.Since(s.startTime),
	}
	if s.startErr != nil {
		res.Err = s.startErr
		return res
	}
	if s.waitErr != nil {
		if errors.Is(s.waitErr, context.Canceled) || errors.Is(s.waitErr, context.DeadlineExceeded) {
			res.Err = s.waitErr
		}
		if res.Err == nil && s.def.ctx != nil && s.def.ctx.Err() != nil {
			res.Err = s.def.ctx.Err()
		}
	}
	if s.cmd.ProcessState != nil {
		res.ExitCode = s.cmd.ProcessState.ExitCode()
		res.signal = signalFromState(s.cmd.ProcessState)
	}
	return res
}

func (c *Cmd) pipelineStages() []*stage {
	root := c.rootCmd()
	stages := []*stage{}
	for current := root; current != nil; current = current.next {
		stages = append(stages, &stage{def: current})
	}
	return stages
}

func (c *Cmd) execCmd() *exec.Cmd {
	cmd := exec.CommandContext(c.ctxOrBackground(), c.name, c.args...)
	if c.dir != "" {
		cmd.Dir = c.dir
	}
	cmd.Env = buildEnv(c.envMode, c.env)
	return cmd
}

func (c *Cmd) stdoutWriter(buf *bytes.Buffer, withCombined bool, combined *bytes.Buffer) io.Writer {
	writers := []io.Writer{buf}
	if withCombined {
		writers = append(writers, combined)
	}
	if c.onStdout != nil {
		writers = append(writers, &lineWriter{onLine: c.onStdout})
	}
	if len(writers) == 1 {
		return buf
	}
	return io.MultiWriter(writers...)
}

func (c *Cmd) stderrWriter(buf *bytes.Buffer, withCombined bool, combined *bytes.Buffer) io.Writer {
	writers := []io.Writer{buf}
	if withCombined {
		writers = append(writers, combined)
	}
	if c.onStderr != nil {
		writers = append(writers, &lineWriter{onLine: c.onStderr})
	}
	if len(writers) == 1 {
		return buf
	}
	return io.MultiWriter(writers...)
}

type lineWriter struct {
	onLine func(string)
	buf    bytes.Buffer
}

func (l *lineWriter) Write(p []byte) (int, error) {
	if l.onLine == nil {
		return len(p), nil
	}
	for _, b := range p {
		if b == '\n' {
			line := l.buf.String()
			l.buf.Reset()
			line = strings.TrimSuffix(line, "\r")
			l.onLine(line)
			continue
		}
		_ = l.buf.WriteByte(b)
	}
	return len(p), nil
}

func buildEnv(mode envMode, env map[string]string) []string {
	merged := map[string]string{}
	if mode != envOnly {
		for _, entry := range os.Environ() {
			key, val, _ := strings.Cut(entry, "=")
			merged[key] = val
		}
	}
	for key, val := range env {
		merged[key] = val
	}
	keys := make([]string, 0, len(merged))
	for key := range merged {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	list := make([]string, 0, len(keys))
	for _, key := range keys {
		list = append(list, key+"="+merged[key])
	}
	return list
}

func shellEscape(arg string) string {
	if arg == "" {
		return "''"
	}
	needsQuote := strings.ContainsAny(arg, " \t\n\r'\"\\$`!")
	if !needsQuote {
		return arg
	}
	return "'" + strings.ReplaceAll(arg, "'", "'\\''") + "'"
}

// Process represents an asynchronously running command.
type Process struct {
	resultOnce sync.Once
	result     Result
	resultCh   chan Result
}

// Wait waits for the command to complete and returns the result.
func (p *Process) Wait() Result {
	p.resultOnce.Do(func() {
		p.result = <-p.resultCh
	})
	return p.result
}
