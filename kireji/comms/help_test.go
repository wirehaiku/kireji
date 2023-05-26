package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func TestFormatHelp(t *testing.T) {
	// setup
	HelpText["test"] = []string{"Test command.", "test"}

	// success
	help := FormatHelp("test")
	assert.Equal(t, "test:\n  Test command.\n  $ kireji test\n", help)
}

func TestHelpCommand(t *testing.T) {
	// setup
	book := book.New(test.Dire(t), ".txt", 0666)

	// failure - invalid command
	err := HelpCommand(book, []string{"nope"})
	assert.EqualError(t, err, `"nope" is not a valid command`)
}
