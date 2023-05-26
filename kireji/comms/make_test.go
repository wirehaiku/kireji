package comms

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestMakeCommand(t *testing.T) {
	// setup
	book := book.New(test.Dire(t), ".txt", 0666)

	// success
	err := MakeCommand(book, []string{"test"})
	assert.FileExists(t, filepath.Join(book.Dire, "test.txt"))
	assert.NoError(t, err)

	// failure - note exists
	err = MakeCommand(book, []string{"test"})
	assert.EqualError(t, err, `note "test" already exists`)
}
