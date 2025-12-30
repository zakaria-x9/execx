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
	// Terminate kills the process immediately.

	// Example: terminate
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		Start()
	_ = proc.Terminate()
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
