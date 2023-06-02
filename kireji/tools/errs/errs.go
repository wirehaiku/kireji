// Package errs implements error creation functions.
package errs

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ArgsMissing returns an error for insufficient arguments.
func ArgsMissing() error {
	return fmt.Errorf("not enough arguments")
}

// CommandAmbiguous returns an error for an ambiguous Command.
func CommandAmbiguous(name string, comms []string) error {
	join := strings.Join(comms, ", ")
	return fmt.Errorf("%s is ambiguous, did you mean: %s?", name, join)
}

// CommandMissing returns an error for a nonexistent Command.
func CommandMissing(name string) error {
	return fmt.Errorf("%s is not a valid command", name)
}

// EnvMissing returns an error for a nonexistent environment variable.
func EnvMissing(name string) error {
	return fmt.Errorf("environment variable %s not set", name)
}

// ExecFail returns an error for a failed program execution.
func ExecFail(prog string, err error) error {
	return fmt.Errorf("%s failed to run: %w", prog, err)
}

// MoveFail returns an error for a failed file move.
func MoveFail(path string, err error) error {
	base := filepath.Base(path)
	return fmt.Errorf("cannot move %s: %w", base, err)
}

// NoteExists returns an error for an existing Note.
func NoteExists(name string) error {
	return fmt.Errorf("%s already exists", name)
}

// NoteMissing returns an error for a nonexistent Note.
func NoteMissing(name string) error {
	return fmt.Errorf("%s does not exist", name)
}

// ReadFail returns an error for a failed file read.
func ReadFail(path string, err error) error {
	base := filepath.Base(path)
	return fmt.Errorf("cannot read %s: %w", base, err)
}

// WriteFail returns an error for a failed file write.
func WriteFail(path string, err error) error {
	base := filepath.Base(path)
	return fmt.Errorf("cannot write %s: %w", base, err)
}
