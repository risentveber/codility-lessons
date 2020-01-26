package sorting

import (
	"testing"

	"github.com/tj/assert"
)

func TestTriangleSolution(t *testing.T) {
	assert.Equal(t, 1, TriangleSolution([]int{3, 3, 4}), "exact one triangle")
}
