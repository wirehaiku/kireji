// Package comms implements command functions and collections.
package comms

import (
	"sort"
	"strings"

	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
)

// Command is a user-callable command function.
type Command func(*book.Book, []string) error

// Commands is a map of all existing commands.
var Commands = map[string]Command{
	"edit":    EditCommand,
	"find":    FindCommand,
	"help":    HelpCommand,
	"junk":    JunkCommand,
	"list":    ListCommand,
	"make":    MakeCommand,
	"move":    MoveCommand,
	"show":    ShowCommand,
	"version": VersionCommand,
}

// Get returns a Command by name, or an error.
func Get(name string) (Command, error) {
	var comms []string

	for comm := range Commands {
		if strings.HasPrefix(comm, name) {
			comms = append(comms, comm)
		}
	}

	sort.Strings(comms)
	switch len(comms) {
	case 0:
		return nil, errs.CommandMissing(name)
	case 1:
		return Commands[comms[0]], nil
	default:
		return nil, errs.CommandAmbiguous(name, comms)
	}
}
