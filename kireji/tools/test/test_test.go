package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

func TestBufs(t *testing.T) {
	// success
	out := Bufs("test\n")
	line := stio.Read()
	stio.Write("test\n")
	assert.Equal(t, "test\n", line)
	assert.Equal(t, "test\n", out.String())
}

func TestDire(t *testing.T) {
	// success
	dire := Dire(t)
	for base, body := range TestFiles {
		bits, _ := os.ReadFile(filepath.Join(dire, base))
		assert.Equal(t, body, string(bits))
	}
}

func TestFile(t *testing.T) {
	// success
	path := File(t)
	bits, _ := os.ReadFile(path)
	assert.Equal(t, TestFiles["alpha.txt"], string(bits))
}
