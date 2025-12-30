//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Arg appends arguments to the command.

	// Example: add args
	cmd := execx.Command("go", "env").Arg("GOOS")
	out, _ := cmd.Output()
	fmt.Println(out != "")
	// #bool true
}
