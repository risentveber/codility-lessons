package maxsliceproblem

import (
	"testing"

	"github.com/tj/assert"
)

func TestMaxProfitSolution(t *testing.T) {
	assert.Equal(t, 356,
		MaxProfitSolution([]int{23171, 21011, 21123, 21366, 21013, 21367}),
		"default example")
	assert.Equal(t, 0, MaxProfitSolution([]int{1, 1, 1, 1, 1, 1, 1}), "all equals")
	assert.Equal(t, 0, MaxProfitSolution([]int{10, 9, 8, 7, 7, 6, 5}), "decreasing")
}
