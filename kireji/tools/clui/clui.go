// Package clui implements command-line user interface functions.
package clui

import (
	"os"

	"github.com/wirehaiku/kireji/kireji/tools/errs"
)

// Arg returns an indexed element from an argument slice, or a default value.
func Arg(args []string, idx int, dflt string) string {
	if len(args) > idx {
		return args[idx]
	}

	return dflt
}

// ArgErr returns an indexed element from an argument slice, or an error.
func ArgErr(args []string, idx int) (string, error) {
	if len(args) > idx {
		return args[idx], nil
	}

	return "", errs.ArgsMissing()
}

// ArgSlice returns a partial arguments slice from an argument slice.
func ArgSlice(args []string, from int) []string {
	if len(args) > from {
		return args[from:]
	}

	return nil
}

// Env returns an existing environment variable through a function, or an error.
func Env(env string, fun func(string) string) (string, error) {
	val, ok := os.LookupEnv(env)
	if !ok || val == "" {
		return "", errs.EnvMissing(env)
	}

	if fun != nil {
		return fun(val), nil
	}

	return val, nil
}
