package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
)

// MakeCommand creates a new empty Note.
func MakeCommand(book *book.Book, args []string) error {
	name, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	if _, err := book.Create(name); err != nil {
		return err
	}

	return nil
}
