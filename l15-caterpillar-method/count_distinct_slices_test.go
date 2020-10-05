package l15caterpillarmethod

import (
	"testing"

	"github.com/tj/assert"
)

func TestCountDistinctSlicesSolution(t *testing.T) {
	assert.Equal(t, 9,
		CountDistinctSlicesSolution(6, []int{3, 4, 5, 5, 2}), "example")
	assert.Equal(t, 6,
		CountDistinctSlicesSolution(3, []int{1, 2, 3}), "simple")
	assert.Equal(t, 7,
		CountDistinctSlicesSolution(3, []int{1, 1, 2, 3}), "simple")
}
