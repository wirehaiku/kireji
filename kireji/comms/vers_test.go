package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestVersionCommand(t *testing.T) {
	// setup
	out := test.Bufs("")
	book := book.New(test.Dire(t), ".txt", 0666)
	Version = "x.y.z (YYYY-MM-DD)"

	// success
	err := VersionCommand(book, []string(nil))
	assert.Equal(t, "Kireji version x.y.z (YYYY-MM-DD).\n", out.String())
	assert.NoError(t, err)
}
