//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// PipeStrict sets strict pipeline semantics.

	// Example: strict
	if len(os.Args) > 2 && os.Args[1] == "execx-example" {
		switch os.Args[2] {
		case "fail":
			os.Exit(2)
		case "ok":
			fmt.Print("ok")
		}
		return
	}
	res := execx.Command(os.Args[0], "execx-example", "fail").
		Pipe(os.Args[0], "execx-example", "ok").
		PipeStrict().
		Run()
	fmt.Println(res.ExitCode)
	// #int 2
}
