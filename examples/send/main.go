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
	// Send sends a signal to the process.

	// Example: send signal
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		Start()
	_ = proc.Send(os.Interrupt)
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
