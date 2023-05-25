package book

import (
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/items/note"
	"github.com/wirehaiku/kireji/kireji/tools/test"
)

func testBook(t *testing.T) *Book {
	dire := test.Dire(t)
	return New(dire, ".txt", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := testBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".txt", book.Extn)
	assert.Equal(t, fs.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	note, err := book.Create("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.txt"), note.Path)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)

	// failure - existing Note
	note, err = book.Create("alpha")
	assert.Nil(t, note)
	assert.EqualError(t, err, `note "alpha" already exists`)
}

func TestGet(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, filepath.Join(book.Dire, "alpha.txt"), note.Path)
	assert.NoError(t, err)

	// failure - nonexistent note
	note, err = book.Get("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `note "nope" does not exist`)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := testBook(t)

	// success - new note
	note, err := book.GetOrCreate("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.txt"), note.Path)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)

	// success - existing note
	note, err = book.GetOrCreate("test")
	assert.Equal(t, filepath.Join(book.Dire, "test.txt"), note.Path)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})
	assert.Len(t, notes, 1)
	assert.Equal(t, filepath.Join(book.Dire, "alpha.txt"), notes[0].Path)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, filepath.Join(book.Dire, "alpha.txt"), notes[0].Path)
	assert.Equal(t, filepath.Join(book.Dire, "bravo.txt"), notes[1].Path)
}

func TestMatch(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	notes := book.Match("alph")
	assert.Len(t, notes, 1)
	assert.Equal(t, filepath.Join(book.Dire, "alpha.txt"), notes[0].Path)
}

func TestSearch(t *testing.T) {
	// setup
	book := testBook(t)

	// success
	notes, err := book.Search("alph")
	assert.Len(t, notes, 1)
	assert.Equal(t, filepath.Join(book.Dire, "alpha.txt"), notes[0].Path)
	assert.NoError(t, err)
}
