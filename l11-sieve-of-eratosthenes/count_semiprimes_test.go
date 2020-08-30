package l11sieveoferatosthenes

import (
	"testing"

	"github.com/tj/assert"
)

func TestCountSemiprimesSolution(t *testing.T) {
	assert.Equal(t, []int{10, 4, 0},
		CountSemiprimesSolution(26, []int{1, 4, 16}, []int{26, 10, 20}), "example")
}
