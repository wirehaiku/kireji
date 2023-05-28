package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
)

func TestGet(t *testing.T) {
	// setup
	Commands["test"] = func(*book.Book, []string) error { return nil }

	// success
	comm, err := Get("test")
	assert.NotNil(t, comm)
	assert.NoError(t, err)

	// failure - invalid command
	comm, err = Get("nope")
	assert.Nil(t, comm)
	assert.Equal(t, errs.CommandMissing("nope"), err)
}
