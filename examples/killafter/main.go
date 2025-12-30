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
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "sleep" {
		time.Sleep(2 * time.Second)
		return
	}
	proc := execx.Command(os.Args[0], "execx-example", "sleep").
		Start()
	proc.KillAfter(100 * time.Millisecond)
	res := proc.Wait()
	fmt.Println(res.ExitCode != 0)
	// #bool true
}
