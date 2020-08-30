package l10primenumbers

import (
	"testing"

	"github.com/tj/assert"
)

func TestPeaksSolution(t *testing.T) {
	assert.Equal(t, 0, PeaksSolution([]int{1, 1, 1, 1, 1, 1, 1}), "no peaks")
	assert.Equal(t, 1, PeaksSolution([]int{1, 1, 1, 2, 1, 1, 1}), "one peak")
	assert.Equal(t, 3, PeaksSolution([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}), "example")
}
