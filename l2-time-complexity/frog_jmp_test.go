package timecomplexity

import (
	"testing"

	"github.com/tj/assert"
)

func TestFrogJmp(t *testing.T) {
	assert.Equal(t, 0, FrogJmpSolution(5, 5, 1), "zero distance")
	assert.Equal(t, 3, FrogJmpSolution(0, 5, 2), "has mod")
	assert.Equal(t, 2, FrogJmpSolution(0, 4, 2), "mod is zero")
}
