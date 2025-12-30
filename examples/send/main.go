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
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "sleep" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0], "execx-example", "sleep").
		Start()
	_ = proc.Send(os.Interrupt)
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
