package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// FindCommand prints all existing Notes containing a substring.
func FindCommand(book *book.Book, args []string) error {
	sub, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	notes, err := book.Search(sub)
	if err != nil {
		return err
	}

	for _, note := range notes {
		stio.Write("%s\n", note.Name())
	}

	return nil
}
