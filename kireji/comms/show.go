package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// ShowCommand prints the contents of an existing note.
func ShowCommand(book *book.Book, args []string) error {
	name, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	note, err := book.Get(name)
	if err != nil {
		return err
	}

	body, err := note.Read()
	if err != nil {
		return err
	}

	stio.Write(body)
	return nil
}
