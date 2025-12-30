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
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "stderr" {
		_, _ = os.Stderr.WriteString("err\n")
		return
	}
	var out strings.Builder
	execx.Command(os.Args[0], "execx-example", "stderr").
		StderrWriter(&out).
		Run()
	fmt.Println(out.Len() > 0)
	// #bool true
}
