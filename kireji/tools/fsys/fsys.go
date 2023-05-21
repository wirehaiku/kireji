// Package fsys implements filesystem handling functions.
package fsys

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Glob returns all paths in a directory matching an extension in sorted order.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	paths, _ := filepath.Glob(glob)
	sort.Strings(paths)
	return paths
}

// Join returns a path from a directory, base name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a file's base name contains a substring.
func Match(path, sub string) bool {
	name := strings.ToLower(Name(path))
	return strings.Contains(name, strings.ToLower(sub))
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

// Search returns true if a file's contents contain a substring.
func Search(path, sub string) (bool, error) {
	body, err := Read(path)
	if err != nil {
		return false, err
	}

	body = strings.ToLower(body)
	return strings.Contains(body, strings.ToLower(sub)), nil
}

// Write writes the contents of a file from a string.
func Write(path, body string, mode fs.FileMode) error {
	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot write file %q: %w", path, err)
	}

	return nil
}
