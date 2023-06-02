package comms

import (
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// Version is the current Kireji version.
var Version = "0.2.0 (2023-05-27)"

// VersionCommand prints the current Kireji version.
func VersionCommand(book *book.Book, args []string) error {
	stio.Write("Kireji version %s.\n", Version)
	return nil
}
