//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// OnStdout registers a line callback for stdout.

	// Example: stdout lines
	var lines []string
	execx.Command("go", "env", "GOOS").
		OnStdout(func(line string) { lines = append(lines, line) }).
		Run()
	fmt.Println(len(lines) > 0)
	// #bool true
}
