package comms

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestEditCommand(t *testing.T) {
	// setup
	dire := t.TempDir()
	os.Setenv("EDITOR", "cat")
	book := book.New(test.Dire(t), ".txt", 0666)
	file, _ := os.Create(filepath.Join(dire, "stdout.txt"))
	os.Stdout = file
	defer file.Close()

	// success
	err := EditCommand(book, []string{"alpha"})
	bits, _ := os.ReadFile(file.Name())
	assert.Equal(t, "Alpha note.\n", string(bits))
	assert.NoError(t, err)
}
