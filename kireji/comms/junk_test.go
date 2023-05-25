package comms

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestJunkCommand(t *testing.T) {
	// setup
	book := book.New(test.Dire(t), ".txt", 0666)

	// success
	err := JunkCommand(book, []string{"alpha"})
	assert.NoFileExists(t, filepath.Join(book.Dire, "alpha.txt"))
	assert.FileExists(t, filepath.Join(book.Dire, "alpha.junk"))
	assert.NoError(t, err)
}
