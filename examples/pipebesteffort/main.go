//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// PipeBestEffort sets best-effort pipeline semantics.

	// Example: best effort
	cmd := execx.Command("go", "env", "GOOS").PipeBestEffort()
	fmt.Println(cmd != nil)
	// #bool true
}
