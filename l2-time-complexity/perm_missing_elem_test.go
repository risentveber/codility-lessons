package timecomplexity

import (
	"testing"

	"github.com/tj/assert"
)

func TestPermMissingElem(t *testing.T) {
	assert.Equal(t, 3, PermMissingElem([]int{2, 4, 1, 5}), "standard example")
	maxValue := 100000
	bigArray := make([]int, maxValue)
	for i := 0; i < maxValue; i++ {
		bigArray[i] = i + 2
	}
	assert.Equal(t, 1, PermMissingElem(bigArray), "max array length")
}
