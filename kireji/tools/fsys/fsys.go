// Package fsys implements filesystem handling functions.
package fsys

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Glob returns all paths in a directory matching an extension.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	paths, _ := filepath.Glob(glob)
	return paths
}

// Name returns a file's base name without the extension.
func Name(path string) string {
	base := filepath.Base(path)
	return strings.Replace(base, filepath.Ext(path), "", 1)
}

// Read returns the contents of a file as a string.
func Read(path string) (string, error) {
	bits, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q: %w", path, err)
	}

	return string(bits), nil
}

// Reext renames a file to a different extension and returns the new path.
func Reext(path, extn string) (string, error) {
	dire := filepath.Dir(path)
	name := Name(path)
	dest := filepath.Join(dire, name+extn)
	if err := os.Rename(path, dest); err != nil {
		return "", fmt.Errorf("cannot rename file %q: %w", path, err)
	}

	return dest, nil
}

// Rename renames a file to a different base name and returns the new path.
func Rename(path, name string) (string, error) {
	dire := filepath.Dir(path)
	extn := filepath.Ext(path)
	dest := filepath.Join(dire, name+extn)
	if err := os.Rename(path, dest); err != nil {
		return "", fmt.Errorf("cannot rename file %q: %w", path, err)
	}

	return dest, nil
}

// Write writes the contents of a file from a string.
func Write(path, body string, mode fs.FileMode) error {
	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot write file %q: %w", path, err)
	}

	return nil
}
