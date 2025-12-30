//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Setsid sets the session ID behavior.

	// Example: setsid
	cmd := execx.Command("go", "env", "GOOS").Setsid(true)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: setsid
	cmd := execx.Command("go", "env", "GOOS").Setsid(true)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: setsid
	cmd := execx.Command("go", "env", "GOOS").Setsid(true)
	fmt.Println(cmd != nil)
	// #bool true
}
