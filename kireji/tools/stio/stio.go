// Package stio implements standard input/output functions.
package stio

import (
	"bufio"
	"fmt"
	"os"
)

// Stdin is the default user input Reader.
var Stdin = bufio.NewReader(os.Stdin)

// Stdout is the default user output Writer.
var Stdout = bufio.NewWriter(os.Stdout)

// Read returns a newline-terminated input string.
func Read() string {
	line, _ := Stdin.ReadString('\n')
	return line
}

// Write writes a formatted output string.
func Write(str string, vars ...any) {
	fmt.Fprintf(Stdout, str, vars...)
	Stdout.Flush()
}
