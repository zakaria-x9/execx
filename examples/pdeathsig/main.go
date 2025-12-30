//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// Pdeathsig is a no-op on non-Linux Unix platforms.

	// Example: pdeathsig
	cmd := execx.Command("go", "env", "GOOS").Pdeathsig(0)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: pdeathsig
	cmd := execx.Command("go", "env", "GOOS").Pdeathsig(0)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: pdeathsig
	cmd := execx.Command("go", "env", "GOOS").Pdeathsig(0)
	fmt.Println(cmd != nil)
	// #bool true
}
