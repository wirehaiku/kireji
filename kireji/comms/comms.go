// Package comms implements command functions and collections.
package comms

import (
	"fmt"

	"github.com/wirehaiku/kireji/kireji/items/book"
)

// Command is a user-callable command function.
type Command func(*book.Book, []string) error

// Commands is a map of all existing commands.
var Commands = map[string]Command{
	"edit": EditCommand,
	"show": ShowCommand,
}

// Get returns a Command by name, or an error.
func Get(name string) (Command, error) {
	comm, ok := Commands[name]
	if !ok {
		return nil, fmt.Errorf("%q is not a valid command", name)
	}

	return comm, nil
}
