package specific

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPackage(t *testing.T) {
	p, err := findPackage("container/ring")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(p.Dir))
	assert.Equal(t, []string{"ring.go"}, p.GoFiles)
	assert.Equal(t, []string{"ring_test.go"}, p.TestGoFiles)
}
