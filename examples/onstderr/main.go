//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// OnStderr registers a line callback for stderr.

	// Example: stderr lines
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		_, _ = os.Stderr.WriteString("err\n")
		return
	}
	var lines []string
	execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		OnStderr(func(line string) { lines = append(lines, line) }).
		Run()
	fmt.Println(len(lines) == 1)
	// #bool true
}
