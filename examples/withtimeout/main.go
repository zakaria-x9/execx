//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"time"
)

func main() {
	// WithTimeout binds the command to a timeout.

	// Example: with timeout
	res := execx.Command("go", "env", "GOOS").WithTimeout(2 * time.Second).Run()
	fmt.Println(res.Err == nil)
	// #bool true
}
