//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// HideWindow controls window visibility.

	// Example: hide window
	cmd := execx.Command("go", "env", "GOOS").HideWindow(true)
	fmt.Println(cmd != nil)
	// #bool true
	// Example: hide window
	cmd := execx.Command("go", "env", "GOOS").HideWindow(true)
	fmt.Println(cmd != nil)
	// #bool true
}
