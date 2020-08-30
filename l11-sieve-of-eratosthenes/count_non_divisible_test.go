package l11sieveoferatosthenes

import (
	"testing"

	"github.com/tj/assert"
)

func TestCountNotDivisibleSolution(t *testing.T) {
	assert.Equal(t, []int{2, 4, 3, 2, 0},
		CountNotDivisibleSolution([]int{3, 1, 2, 3, 6}), "example")
}
