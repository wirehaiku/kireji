package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// EditCommand opens an existing note in the default editor.
func EditCommand(book *book.Book, args []string) error {
	name, err := clui.ArgErr(args, 0)
	if err != nil {
		return err
	}

	note, err := book.GetOrCreate(name)
	if err != nil {
		return err
	}

	edit, err := clui.Env("EDITOR", nil)
	if err != nil {
		return err
	}

	if err := stio.Exec(edit, note.Path); err != nil {
		return err
	}

	return nil
}
