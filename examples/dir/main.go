//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// Dir sets the working directory.

	// Example: change dir
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		wd, _ := os.Getwd()
		fmt.Println(wd)
		return
	}
	dir := os.TempDir()
	out, _ := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		Dir(dir).
		OutputTrimmed()
	fmt.Println(out == dir)
	// #bool true
}
