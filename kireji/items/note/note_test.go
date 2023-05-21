package note

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func testNote(t *testing.T) *Note {
	path := test.File(t)
	return New(path, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := testNote(t)
	assert.Contains(t, note.Path, "alpha.txt")
	assert.Equal(t, fs.FileMode(0666), note.Mode)
}

func TestMatch(t *testing.T) {
	// setup
	note := testNote(t)

	// success
	ok := note.Match("alph")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := testNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := testNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	note := testNote(t)
	orig := note.Path

	// success
	err := note.Rename("test")
	assert.Contains(t, note.Path, "test.txt")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)
}

func TestTrash(t *testing.T) {
	// setup
	note := testNote(t)
	orig := note.Path

	// success
	err := note.Trash()
	assert.Contains(t, note.Path, "alpha.trash")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := testNote(t)

	// success
	ok, err := note.Search("alph")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	note := testNote(t)

	// success
	err := note.Write("test\n")
	bits, _ := os.ReadFile(note.Path)
	assert.Equal(t, "test\n", string(bits))
	assert.NoError(t, err)
}
