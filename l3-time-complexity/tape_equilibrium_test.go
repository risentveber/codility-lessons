package timecomplexity

import (
	"testing"

	"github.com/tj/assert"
)

func TestTapeEquilibrium(t *testing.T) {
	assert.Equal(t, 1, TapeEquilibriumSolution([]int{3, 1, 2, 4, 3}), "standard example")
}
