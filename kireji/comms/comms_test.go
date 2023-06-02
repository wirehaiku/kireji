package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
)

func TestGet(t *testing.T) {
	// setup
	Commands = map[string]Command{
		"comm1": func(*book.Book, []string) error { return nil },
		"comm2": func(*book.Book, []string) error { return nil },
		"test":  func(*book.Book, []string) error { return nil },
	}

	// success
	comm, err := Get("t")
	assert.NotNil(t, comm)
	assert.NoError(t, err)

	// failure - missing command
	comm, err = Get("nope")
	assert.Nil(t, comm)
	assert.Equal(t, errs.CommandMissing("nope"), err)

	// failure - ambiguous command
	comm, err = Get("comm")
	assert.Nil(t, comm)
	assert.Equal(t, errs.CommandAmbiguous("comm", []string{"comm1", "comm2"}), err)
}
