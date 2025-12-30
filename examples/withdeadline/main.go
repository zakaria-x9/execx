//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"time"
)

func main() {
	// WithDeadline binds the command to a deadline.

	// Example: with deadline
	res := execx.Command("go", "env", "GOOS").WithDeadline(time.Now().Add(2 * time.Second)).Run()
	fmt.Println(res.Err == nil)
	// #bool true
}
