//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
)

func main() {
	// PipelineResults executes the command and returns per-stage results.

	// Example: pipeline results
	results := execx.Command("go", "env", "GOOS").PipelineResults()
	fmt.Println(len(results) == 1)
	// #bool true
}
