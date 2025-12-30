//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Command constructs a new command without executing it.

	// Example: command
	cmd := execx.Command("go", "env", "GOOS")
	out, _ := cmd.Output()
	fmt.Println(out != "")
	// #bool true
}
