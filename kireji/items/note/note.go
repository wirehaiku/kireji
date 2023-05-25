// Package note implements the Note type and methods.
package note

import (
	"io/fs"

	"github.com/wirehaiku/kireji/kireji/tools/fsys"
	"github.com/wirehaiku/kireji/kireji/tools/neat"
)

// Note is a single plaintext file in a directory.
type Note struct {
	Path string
	Mode fs.FileMode
}

// New returns a new Note.
func New(path string, mode fs.FileMode) *Note {
	return &Note{neat.Path(path), mode}
}

// Exists returns true if the Note's file exists.
func (n *Note) Exists() bool {
	return fsys.Exists(n.Path)
}

// Junk renames the Note's file to the ".junk" extension.
func (n *Note) Junk() error {
	dest, err := fsys.Reext(n.Path, ".junk")
	if err != nil {
		return err
	}

	n.Path = dest
	return nil
}

// Match returns true if the Note's name contains a substring.
func (n *Note) Match(sub string) bool {
	return fsys.Match(n.Path, sub)
}

// Name returns the Note's base name.
func (n *Note) Name() string {
	return neat.Name(fsys.Name(n.Path))
}

// Read returns the Note's contents as a string.
func (n *Note) Read() (string, error) {
	body, err := fsys.Read(n.Path)
	return neat.Body(body), err
}

// Rename renames the Note's file.
func (n *Note) Rename(name string) error {
	dest, err := fsys.Rename(n.Path, neat.Name(name))
	if err != nil {
		return err
	}

	n.Path = dest
	return nil
}

// Search returns true if the Note's contents contain a substring.
func (n *Note) Search(sub string) (bool, error) {
	return fsys.Search(n.Path, sub)
}

// Write writes the Note's contents from a string.
func (n *Note) Write(body string) error {
	return fsys.Write(n.Path, neat.Body(body), n.Mode)
}
