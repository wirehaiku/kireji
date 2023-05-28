package clui

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kireji/kireji/tools/errs"
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

func TestArgSlice(t *testing.T) {
	// success - partial slice
	args := ArgSlice([]string{"one", "two"}, 1)
	assert.Equal(t, []string{"two"}, args)

	// success - empty slice
	args = ArgSlice([]string{"one", "two"}, 2)
	assert.Empty(t, args)
}

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("EMPTY", "")
	os.Setenv("TEST", "test")

	// success - with function
	val, err := Env("TEST", strings.ToUpper)
	assert.Equal(t, "TEST", val)
	assert.NoError(t, err)

	// success - nil function
	val, err = Env("TEST", nil)
	assert.Equal(t, "test", val)
	assert.NoError(t, err)

	// failure - missing env
	val, err = Env("NOPE", strings.ToUpper)
	assert.Empty(t, val)
	assert.Equal(t, errs.EnvMissing("NOPE"), err)

	// failure - empty env
	val, err = Env("EMPTY", strings.ToUpper)
	assert.Empty(t, val)
	assert.Equal(t, errs.EnvMissing("EMPTY"), err)
}
