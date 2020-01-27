package stacksandqueues

import (
	"testing"

	"github.com/tj/assert"
)

func TestFishSolution(t *testing.T) {
	assert.Equal(t, 4, FishSolution([]int{2, 3, 4, 5}, []int{0, 0, 1, 1}), "all in opposite")
	assert.Equal(t, 2, FishSolution([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}), "all in opposite")
}
