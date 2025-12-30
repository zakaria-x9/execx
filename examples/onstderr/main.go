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
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "stderr" {
		_, _ = os.Stderr.WriteString("err\n")
		return
	}
	var lines []string
	execx.Command(os.Args[0], "execx-example", "stderr").
		OnStderr(func(line string) { lines = append(lines, line) }).
		Run()
	fmt.Println(len(lines) == 1)
	// #bool true
}
