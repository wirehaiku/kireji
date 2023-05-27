package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestFindCommand(t *testing.T) {
	// setup
	out := test.Bufs("")
	book := book.New(test.Dire(t), ".txt", 0666)

	// success
	err := FindCommand(book, []string{"alph"})
	assert.Equal(t, "alpha\n", out.String())
	assert.NoError(t, err)
}
