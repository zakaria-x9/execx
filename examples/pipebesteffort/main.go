//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
	"time"
)

func main() {
	// PipeBestEffort sets best-effort pipeline semantics.

	// Example: best effort
	if len(os.Args) > 2 && os.Args[1] == "execx-example" {
		switch os.Args[2] {
		case "sleep":
			time.Sleep(200 * time.Millisecond)
		case "ok":
			fmt.Print("ok")
		}
		return
	}
	res := execx.Command(os.Args[0], "execx-example", "sleep").
		WithTimeout(50 * time.Millisecond).
		Pipe(os.Args[0], "execx-example", "ok").
		PipeBestEffort().
		Run()
	fmt.Println(res.Stdout)
	// #string ok
}
