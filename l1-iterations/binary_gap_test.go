package iterations

import (
	"testing"

	"github.com/tj/assert"
)

func TestContainerCreation(t *testing.T) {
	assert.Equal(t, 0, BinaryGapSolution(0), "0")
	assert.Equal(t, 5, BinaryGapSolution(1041), "1041")
	assert.Equal(t, 0, BinaryGapSolution(15), "15")
	assert.Equal(t, 0, BinaryGapSolution(32), "32")
}
