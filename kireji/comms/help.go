package comms

import (
	"fmt"
	"strings"

	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
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

	"move": {
		"Move an existing note to a new name.",
		"move NOTE NAME",
	},

	"show": {
		"Print the contents of a note.",
		"show NOTE",
	},

	"version": {
		"Print the current version.",
		"version",
	},
}

// FormatHelp formats a Command's help message.
func FormatHelp(name string) string {
	return strings.Join([]string{
		fmt.Sprintf("%s:", name),
		fmt.Sprintf("  %s", HelpText[name][0]),
		fmt.Sprintf("  $ kireji %s", HelpText[name][1]),
	}, "\n") + "\n"
}

// HelpCommand prints help for one or all existing Commands.
func HelpCommand(book *book.Book, args []string) error {
	name := clui.Arg(args, 0, "")
	_, ok := HelpText[name]

	switch {
	case ok:
		stio.Write(FormatHelp(name) + "\n")
		return nil

	case name == "":
		stio.Write("Kireji: A plaintext note-taking CLI in Go.\n\n")
		for name := range HelpText {
			stio.Write(FormatHelp(name) + "\n")
		}

		stio.Write("See github.com/wirehaiku/kireji for help & bugs.\n")
		return nil

	default:
		return errs.CommandMissing(name)
	}
}
