package stio

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
