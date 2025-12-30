//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
)

func main() {
	// StdinFile sets stdin from a file.

	// Example: stdin file
	if len(os.Args) > 2 && os.Args[1] == "execx-example" && os.Args[2] == "stdin" {
		buf := make([]byte, 8)
		n, _ := os.Stdin.Read(buf)
		_, _ = os.Stdout.Write(buf[:n])
		return
	}
	file, _ := os.CreateTemp("", "execx-stdin")
	_, _ = file.WriteString("hi")
	_, _ = file.Seek(0, 0)
	out, _ := execx.Command(os.Args[0], "execx-example", "stdin").
		StdinFile(file).
		Output()
	fmt.Println(out == "hi")
	// #bool true
}
