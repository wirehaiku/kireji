// Package book implements the Book type and methods.
package book

import (
	"io/fs"

	"github.com/wirehaiku/kireji/kireji/items/note"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
	"github.com/wirehaiku/kireji/kireji/tools/fsys"
	"github.com/wirehaiku/kireji/kireji/tools/neat"
)

// Book is a directory of plaintext Note files.
type Book struct {
	Dire string
	Extn string
	Mode fs.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode fs.FileMode) *Book {
	return &Book{neat.Path(dire), neat.Extn(extn), mode}
}

// Create creates and returns a new empty Note.
func (b *Book) Create(name string) (*note.Note, error) {
	path := fsys.Join(b.Dire, neat.Name(name), b.Extn)
	note := note.New(path, b.Mode)
	if note.Exists() {
		return nil, errs.NoteExists(name)
	}

	if err := note.Write(""); err != nil {
		return nil, err
	}

	return note, nil
}

// Get returns an existing Note by name.
func (b *Book) Get(name string) (*note.Note, error) {
	dest := fsys.Join(b.Dire, neat.Name(name), b.Extn)
	note := note.New(dest, b.Mode)
	if !note.Exists() {
		return nil, errs.NoteMissing(name)
	}

	return note, nil
}

// GetOrCreate returns an existing or a newly created Note.
func (b *Book) GetOrCreate(name string) (*note.Note, error) {
	dest := fsys.Join(b.Dire, neat.Name(name), b.Extn)
	note := note.New(dest, b.Mode)
	if note.Exists() {
		return note, nil
	}

	return b.Create(name)
}

// Filter returns all Notes passing a filter function.
func (b *Book) Filter(fun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note
	for _, note := range b.List() {
		ok, err := fun(note)
		if err != nil {
			return nil, err
		}

		if ok {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// List returns all Notes in the Book's directory in sorted order.
func (b *Book) List() []*note.Note {
	var notes []*note.Note
	for _, path := range fsys.Glob(b.Dire, b.Extn) {
		notes = append(notes, note.New(path, b.Mode))
	}

	return notes
}

// Match returns all Notes with names containing a substring.
func (b *Book) Match(sub string) []*note.Note {
	notes, _ := b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(sub), nil
	})

	return notes
}

// Search returns all Notes with bodies containing a substring.
func (b *Book) Search(sub string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(sub)
	})
}
