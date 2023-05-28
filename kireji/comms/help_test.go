package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatHelp(t *testing.T) {
	// setup
	HelpText["test"] = []string{"Test command.", "test"}

	// success
	help := FormatHelp("test")
	assert.Equal(t, "test:\n  Test command.\n  $ kireji test\n", help)
}

func TestHelpCommand(t *testing.T) {
	// no need to test command
}
