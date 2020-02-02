package leader

import (
	"testing"

	"github.com/tj/assert"
)

func TestDominatorSolution(t *testing.T) {
	data := []int{0, 0, 1}
	result := DominatorSolution(data)
	assert.Equal(t, 0, data[result], "value is 0")
	assert.Equal(t, -1, DominatorSolution([]int{0, 1}), "no dominator")
	assert.Equal(t, -1, DominatorSolution([]int{}), "no dominator in empty array")
	assert.Equal(t, -1, DominatorSolution([]int{0, 1, 2}), "no dominator")
}
