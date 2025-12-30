//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// StdinBytes sets stdin from bytes.

	// Example: stdin bytes
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "stdin" {
		buf := make([]byte, 8)
		n, _ := os.Stdin.Read(buf)
		_, _ = os.Stdout.Write(buf[:n])
		return
	}
	out, _ := execx.Command(os.Args[0], "execx-example", "stdin").
		StdinBytes([]byte("hi")).
		Output()
	fmt.Println(out == "hi")
	// #bool true
}
