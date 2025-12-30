//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// CreationFlags sets Windows creation flags.

	// Example: creation flags
	cmd := execx.Command("go", "env", "GOOS").CreationFlags(0)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: creation flags
	cmd := execx.Command("go", "env", "GOOS").CreationFlags(0)
	fmt.Println(cmd != nil)
	// #bool true
}
