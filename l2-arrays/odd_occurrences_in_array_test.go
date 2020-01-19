package arrays

import (
	"testing"

	"github.com/tj/assert"
)

func TestOddOccurrencesInArray(t *testing.T) {
	assert.Equal(t, 1, OddOccurrencesInArraySolution([]int{1}), "only one element")
	assert.Equal(t, 2, OddOccurrencesInArraySolution([]int{1, 2, 3, 1, 3}), "default")
}
