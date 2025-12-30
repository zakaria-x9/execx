//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Setpgid sets the process group ID behavior.

	// Example: setpgid
	cmd := execx.Command("go", "env", "GOOS").Setpgid(true)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: setpgid
	cmd := execx.Command("go", "env", "GOOS").Setpgid(true)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: setpgid
	cmd := execx.Command("go", "env", "GOOS").Setpgid(true)
	fmt.Println(cmd != nil)
	// #bool true
}
