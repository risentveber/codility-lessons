package leader

import (
	"testing"

	"github.com/tj/assert"
)

func TestEquiLeaderSolution(t *testing.T) {
	assert.Equal(t, 0, EquiLeaderSolution([]int{0, 1}), "no dominator")
	assert.Equal(t, 0, EquiLeaderSolution([]int{}), "no dominator in empty array")
	assert.Equal(t, 2, EquiLeaderSolution([]int{4, 3, 4, 4, 4, 2}), "default example")
}
