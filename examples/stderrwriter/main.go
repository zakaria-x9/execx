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
	// StderrWriter sets a raw writer for stderr.

	// Example: stderr writer
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		_, _ = os.Stderr.WriteString("err\n")
		return
	}
	var out strings.Builder
	execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		StderrWriter(&out).
		Run()
	fmt.Println(out.Len() > 0)
	// #bool true
}
