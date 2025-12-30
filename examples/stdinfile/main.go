//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// StdinFile sets stdin from a file.

	// Example: stdin file
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		buf := make([]byte, 8)
		n, _ := os.Stdin.Read(buf)
		_, _ = os.Stdout.Write(buf[:n])
		return
	}
	file, _ := os.CreateTemp("", "execx-stdin")
	_, _ = file.WriteString("hi")
	_, _ = file.Seek(0, 0)
	out, _ := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		StdinFile(file).
		Output()
	fmt.Println(out == "hi")
	// #bool true
}
