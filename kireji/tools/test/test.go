// Package test implements unit-testing data and helper functions.
package test

import (
	"os"
	"path/filepath"
	"testing"
)

var TestFiles = map[string]string{
	"alpha.txt":     "Alpha note.\n",
	"bravo.txt":     "Bravo note.\n",
	"charlie.trash": "Charlie note (trashed).\n",
}

// Dire returns a temporary directory populated with TestFiles.
func Dire(t *testing.T) string {
	dire := t.TempDir()

	for base, body := range TestFiles {
		os.WriteFile(filepath.Join(dire, base), []byte(body), 0666)
	}

	return dire
}

// File returns the "alpha.txt" file from the Dire temporary directory.
func File(t *testing.T) string {
	dire := Dire(t)
	return filepath.Join(dire, "alpha.txt")
}
