package comms

import (
	"fmt"
	"strings"

	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// HelpText is a map of help messages for all existing Commands.
var HelpText = map[string][]string{
	"edit": {
		"Edit a new or existing note.",
		"edit NOTE",
	},

	"help": {
		"Show help for one or all Kireji commands.",
		"help [COMMAND]",
	},

	"junk": {
		"Junk an existing note.",
		"junk NOTE",
	},

	"list": {
		"List all notes, or notes matching SUBSTRING.",
		"list [SUBSTRING]",
	},

	"make": {
		"Create a new empty note.",
		"make NOTE",
	},

	"show": {
		"Print the contents of a note.",
		"show NOTE",
	},
}

// FormatHelp formats a Command's help message.
func FormatHelp(comm string) string {
	return strings.Join([]string{
		fmt.Sprintf("%s:", comm),
		fmt.Sprintf("  %s", HelpText[comm][0]),
		fmt.Sprintf("  $ kireji %s", HelpText[comm][1]),
	}, "\n") + "\n"
}

// HelpCommand prints help for one or all existing Commands.
func HelpCommand(book *book.Book, args []string) error {
	comm := clui.Arg(args, 0, "")
	_, ok := HelpText[comm]

	switch {
	case ok:
		stio.Write(FormatHelp(comm) + "\n")
		return nil

	case comm == "":
		stio.Write("Kireji: A plaintext note-taking CLI in Go.\n\n")
		for comm := range HelpText {
			stio.Write(FormatHelp(comm) + "\n")
		}

		stio.Write("See github.com/wirehaiku/kireji for help & bugs.\n")
		return nil

	default:
		return fmt.Errorf("%q is not a valid command", comm)
	}
}
