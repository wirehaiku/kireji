package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
)

// JunkCommand junks an existing Note.
func JunkCommand(book *book.Book, args []string) error {
	name, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	note, err := book.Get(name)
	if err != nil {
		return err
	}

	if err := note.Junk(); err != nil {
		return err
	}

	return nil
}
