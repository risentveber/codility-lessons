package arrays

import (
	"testing"

	"github.com/tj/assert"
)

func TestContainerCreation(t *testing.T) {
	assert.Equal(t, []int{9, 7, 6, 3, 8}, CyclicRotationSolution([]int{3, 8, 9, 7, 6}, 3), "default")
	assert.Equal(t, []int{9, 7, 6, 3, 8}, CyclicRotationSolution([]int{3, 8, 9, 7, 6}, -7), "default negative")
	assert.Equal(t, []int{}, CyclicRotationSolution([]int{}, 2), "empty array")
}
