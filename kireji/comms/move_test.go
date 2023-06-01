package comms

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestMoveCommand(t *testing.T) {
	// setup
	book := book.New(test.Dire(t), ".txt", 0666)

	// success
	err := MoveCommand(book, []string{"alpha", "test"})
	assert.FileExists(t, filepath.Join(book.Dire, "test.txt"))
	assert.NoFileExists(t, filepath.Join(book.Dire, "alpha.txt"))
	assert.NoError(t, err)
}
