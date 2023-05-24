package stio

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	// setup
	dire := t.TempDir()
	file, _ := os.Create(filepath.Join(dire, "stdout.txt"))
	os.Stdout = file
	defer file.Close()

	// success
	err := Exec("echo", "test")
	bits, _ := os.ReadFile(file.Name())
	assert.Equal(t, "test\n", string(bits))
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	// setup
	buf := bytes.NewBufferString("test\n")
	Stdin = bufio.NewReader(buf)

	// success
	line := Read()
	assert.Equal(t, "test\n", line)
}

func TestWrite(t *testing.T) {
	// setup
	buf := bytes.NewBuffer(nil)
	Stdout = bufio.NewWriter(buf)

	// success
	Write("%s\n", "test")
	assert.Equal(t, "test\n", buf.String())
}
