package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
)

// MoveCommand renames an existing Note to a new name.
func MoveCommand(book *book.Book, args []string) error {
	name, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	dest, err := clui.ArgErr(args, 1)
	if err != nil {
		return err
	}

	note, err := book.Get(name)
	if err != nil {
		return err
	}

	return note.Rename(dest)
}
