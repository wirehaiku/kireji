// Kireji: A plaintext note-taking CLI in Go.
package main

import (
	"os"

	"github.com/wirehaiku/kireji/kireji/comms"
	"github.com/wirehaiku/kireji/kireji/items/book"
	"github.com/wirehaiku/kireji/kireji/tools/clui"
	"github.com/wirehaiku/kireji/kireji/tools/neat"
	"github.com/wirehaiku/kireji/kireji/tools/stio"
)

// die fatally logs an error message.
func die(str string) {
	stio.Write("Error: %s.\n", str)
	os.Exit(1)
}

// try fatally logs any non-nil error.
func try(err error) {
	if err != nil {
		die(err.Error())
	}
}

// main runs the main Kireji program.
func main() {
	dire, err := clui.Env("KIREJI_DIR", neat.Path)
	try(err)

	extn, err := clui.Env("KIREJI_EXT", neat.Extn)
	try(err)

	name := clui.Arg(os.Args, 1, "help")
	args := clui.ArgSlice(os.Args, 2)
	book := book.New(dire, extn, 0666)

	comm, err := comms.Get(name)
	try(err)
	try(comm(book, args))
}
