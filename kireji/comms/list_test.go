package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestListCommand(t *testing.T) {
	// setup
	out := test.Bufs("")
	book := book.New(test.Dire(t), ".txt", 0666)

	// success - no argument
	err := ListCommand(book, []string(nil))
	assert.Equal(t, "alpha\nbravo\n", out.String())
	assert.NoError(t, err)

	// setup
	out.Reset()

	// success - with argument
	err = ListCommand(book, []string{"alph"})
	assert.Equal(t, "alpha\n", out.String())
	assert.NoError(t, err)
}
