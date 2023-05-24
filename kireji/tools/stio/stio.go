// Package stio implements standard input/output functions.
package stio

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Stdin is the default user input Reader.
var Stdin = bufio.NewReader(os.Stdin)

// Stdout is the default user output Writer.
var Stdout = bufio.NewWriter(os.Stdout)

// Exec executes an external command with arguments.
func Exec(name string, args ...string) error {
	comm := exec.Command(name, args...)
	comm.Stdin = Stdin
	comm.Stdout = Stdout
	if err := comm.Run(); err != nil {
		return fmt.Errorf("failed to run %q: %w", name, err)
	}

	return nil
}

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
