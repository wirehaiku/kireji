package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// ListCommand prints all existing Notes, or Notes matching a substring.
func ListCommand(book *book.Book, args []string) error {
	sub := clui.Arg(args, 0, "")

	for _, note := range book.Match(sub) {
		stio.Write("%s\n", note.Name())
	}

	return nil
}
