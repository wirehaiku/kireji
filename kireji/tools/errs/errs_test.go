package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsMissing(t *testing.T) {
	// success
	err := ArgsMissing()
	assert.EqualError(t, err, "not enough arguments")
}

func TestCommandMissing(t *testing.T) {
	// success
	err := CommandMissing("test")
	assert.EqualError(t, err, "test is not a valid command")
}

func TestEnvMissing(t *testing.T) {
	// success
	err := EnvMissing("test")
	assert.EqualError(t, err, "environment variable test not set")
}

func TestExecFail(t *testing.T) {
	// success
	err := ExecFail("test", errors.New("error"))
	assert.EqualError(t, err, "test failed to run: error")
}

func TestMoveFail(t *testing.T) {
	// success
	err := MoveFail("/dire/test.txt", errors.New("error"))
	assert.EqualError(t, err, "cannot move test.txt: error")
}

func TestNoteExists(t *testing.T) {
	// success
	err := NoteExists("test")
	assert.EqualError(t, err, "test already exists")
}

func TestNoteMissing(t *testing.T) {
	// success
	err := NoteMissing("test")
	assert.EqualError(t, err, "test does not exist")
}

func TestReadFail(t *testing.T) {
	// success
	err := ReadFail("/dire/test.txt", errors.New("error"))
	assert.EqualError(t, err, "cannot read test.txt: error")
}

func TestWriteFail(t *testing.T) {
	// success
	err := WriteFail("/dire/test.txt", errors.New("error"))
	assert.EqualError(t, err, "cannot write test.txt: error")
}
