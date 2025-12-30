//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// PipeStrict sets strict pipeline semantics.

	// Example: strict
	cmd := execx.Command("go", "env", "GOOS").PipeStrict()
	fmt.Println(cmd != nil)
	// #bool true
}
