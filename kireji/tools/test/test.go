// Package test implements unit-testing data and helper functions.
package test

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

var TestFiles = map[string]string{
	"alpha.txt":    "Alpha note.\n",
	"bravo.txt":    "Bravo note.\n",
	"charlie.junk": "Charlie note (junked).\n",
}

// Bufs sets Stdin and Stdout to temporary buffers, populated Stdin with a string and
// returns Stdout's buffer.
func Bufs(line string) *bytes.Buffer {
	inn := bytes.NewBufferString(line)
	out := bytes.NewBuffer(nil)
	stio.Stdin = bufio.NewReader(inn)
	stio.Stdout = bufio.NewWriter(out)
	return out
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
