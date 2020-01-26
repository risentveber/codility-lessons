package sorting

import (
	"testing"

	"github.com/tj/assert"
)

func TestMaxProductOfThreeSolution(t *testing.T) {
	assert.Equal(t, 36, MaxProductOfThreeSolution([]int{3, 3, 4}), "only one product")
	assert.Equal(t, 60, MaxProductOfThreeSolution([]int{-3, 1, 2, -2, 5, 6}), "from example")
}
