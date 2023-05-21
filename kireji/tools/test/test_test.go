package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
