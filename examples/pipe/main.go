//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// Pipe appends a new command to the pipeline.

	// Example: pipe
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		mode := os.Getenv("EXECX_EXAMPLE_MODE")
		if mode == "emit" {
			fmt.Print("ok")
			return
		}
		if mode == "echo" {
			buf := make([]byte, 8)
			n, _ := os.Stdin.Read(buf)
			_, _ = os.Stdout.Write(buf[:n])
			return
		}
	}
	out, _ := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1", "EXECX_EXAMPLE_MODE=emit").
		Pipe(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1", "EXECX_EXAMPLE_MODE=echo").
		Output()
	fmt.Println(out == "ok")
	// #bool true
}
