package l14binarysearch

import (
	"testing"

	"github.com/tj/assert"
)

func TestMinMaxDivisionSolution(t *testing.T) {
	assert.Equal(t, 6,
		MinMaxDivisionSolution(3, 5, []int{2, 1, 5, 1, 2, 2, 2}), "example")
}
