// Package comms implements command functions and collections.
package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
)

// Command is a user-callable command function.
type Command func(*book.Book, []string) error

// Commands is a map of all existing commands.
var Commands = map[string]Command{
	"edit": EditCommand,
	"find": FindCommand,
	"help": HelpCommand,
	"junk": JunkCommand,
	"list": ListCommand,
	"make": MakeCommand,
	"show": ShowCommand,
}

// Get returns a Command by name, or an error.
func Get(name string) (Command, error) {
	comm, ok := Commands[name]
	if !ok {
		return nil, errs.CommandMissing(name)
	}

	return comm, nil
}
