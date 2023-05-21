package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\nbody\n\n")
	assert.Equal(t, "body\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("\t.EXT\n")
	assert.Equal(t, ".ext", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tFOO BAR\n")
	assert.Equal(t, "foo-bar", name)
}
