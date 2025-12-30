//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/execx"
	"os"
	"strings"
)

func main() {
	// StdinReader sets stdin from an io.Reader.

	// Example: stdin reader
	if os.Getenv("EXECX_EXAMPLE_CHILD") == "1" {
		buf := make([]byte, 8)
		n, _ := os.Stdin.Read(buf)
		_, _ = os.Stdout.Write(buf[:n])
		return
	}
	out, _ := execx.Command(os.Args[0]).
		Env("EXECX_EXAMPLE_CHILD=1").
		StdinReader(strings.NewReader("hi")).
		Output()
	fmt.Println(out == "hi")
	// #bool true
}
