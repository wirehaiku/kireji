package fsys

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestGlob(t *testing.T) {
	// setup
	dire := test.Dire(t)

	// success
	paths := Glob(dire, ".txt")
	assert.Equal(t, []string{
		dire + "/alpha.txt",
		dire + "/bravo.txt",
	}, paths)
}

func TestName(t *testing.T) {
	// success
	name := Name("/dire/name.ext")
	assert.Equal(t, "name", name)
}

func TestRead(t *testing.T) {
	// setup
	path := test.File(t)

	// success
	body, err := Read(path)
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestReext(t *testing.T) {
	// setup
	path := test.File(t)

	// success
	dest, err := Reext(path, ".test")
	assert.Equal(t, filepath.Dir(path)+"/alpha.test", dest)
	assert.NoFileExists(t, filepath.Dir(path)+"/alpha.txt")
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	path := test.File(t)

	// success
	dest, err := Rename(path, "test")
	assert.Equal(t, filepath.Dir(path)+"/test.txt", dest)
	assert.NoFileExists(t, filepath.Dir(path)+"/alpha.txt")
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	path := test.File(t)

	// success
	err := Write(path, "test", 0666)
	bits, _ := os.ReadFile(path)
	assert.Equal(t, "test", string(bits))
	assert.NoError(t, err)
}
