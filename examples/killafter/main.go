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
	// KillAfter terminates the process after the given duration.

	// Example: kill after
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		Start()
	proc.KillAfter(100 * time.Millisecond)
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
