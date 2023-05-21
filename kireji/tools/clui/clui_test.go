package clui

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArg(t *testing.T) {
	// success - indexed value
	arg := Arg([]string{"one", "two"}, 1, "default")
	assert.Equal(t, "two", arg)

	// success - default value
	arg = Arg([]string{"one", "two"}, 2, "default")
	assert.Equal(t, "default", arg)
}

func TestArgErr(t *testing.T) {
	// success
	arg, err := ArgErr([]string{"one", "two"}, 1)
	assert.Equal(t, "two", arg)
	assert.NoError(t, err)

	// failure - out of bounds
	arg, err = ArgErr([]string{"one", "two"}, 2)
	assert.Empty(t, arg)
	assert.EqualError(t, err, "not enough arguments")
}

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("EMPTY", "")
	os.Setenv("TEST", "test")

	// success
	val, err := Env("TEST", strings.ToUpper)
	assert.Equal(t, "TEST", val)

	// failure - missing env
	val, err = Env("NOPE", strings.ToUpper)
	assert.Empty(t, val)
	assert.EqualError(t, err, `environment variable "NOPE" not set`)

	// failure - empty env
	val, err = Env("EMPTY", strings.ToUpper)
	assert.Empty(t, val)
	assert.EqualError(t, err, `environment variable "EMPTY" not set`)
}
