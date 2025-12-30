//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
	"strings"
)

func main() {
	// PipelineResults executes the command and returns per-stage results.

	// Example: pipeline results
	if len(os.Args) > 2 && os.Args[1] == "execx-example" {
		switch os.Args[2] {
		case "emit":
			fmt.Print("go")
		case "upper":
			buf := make([]byte, 8)
			n, _ := os.Stdin.Read(buf)
			fmt.Print(strings.ToUpper(string(buf[:n])))
		}
		return
	}
	results := execx.Command(os.Args[0], "execx-example", "emit").
		Pipe(os.Args[0], "execx-example", "upper").
		PipelineResults()
	fmt.Println(len(results))
	// #int 2
}
