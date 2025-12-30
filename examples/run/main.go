//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Run executes the command and returns the result.

	// Example: run
	res := execx.Command("go", "env", "GOOS").Run()
	fmt.Println(res.ExitCode == 0)
	// #bool true
}
