package maxsliceproblem

import (
	"testing"

	"github.com/tj/assert"
)

func TestNumberOfDiskIntersectionsSolution(t *testing.T) {
	assert.Equal(t, 17, MaxDoubleSliceSolution([]int{3, 2, 6, -1, 4, 5, -1, 2}), "default example")
	assert.Equal(t, 4, MaxDoubleSliceSolution([]int{1, 1, 1, 1, 1, 1, 1}), "all ones")
}
