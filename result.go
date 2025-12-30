package execx

import (
	"os"
	"time"
)

// Result captures the outcome of a command execution.
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Err      error
	Duration time.Duration

	signal os.Signal
}

// OK reports whether the command exited cleanly without errors.
func (r Result) OK() bool {
	return r.Err == nil && r.ExitCode == 0
}

// IsExitCode reports whether the exit code matches.
func (r Result) IsExitCode(code int) bool {
	return r.ExitCode == code
}

// IsSignal reports whether the command terminated due to a signal.
func (r Result) IsSignal(sig os.Signal) bool {
	return r.signal == sig
}
