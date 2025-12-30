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
	// GracefulShutdown sends a signal and escalates to kill after the timeout.

	// Example: graceful shutdown
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "sleep" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0], "execx-example", "sleep").
		Start()
	_ = proc.GracefulShutdown(os.Interrupt, 100*time.Millisecond)
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
