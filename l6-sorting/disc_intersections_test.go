package sorting

import (
	"testing"

	"github.com/tj/assert"
)

func TestNumberOfDiskIntersectionsSolution(t *testing.T) {
	assert.Equal(t, 11, NumberOfDiskIntersectionsSolution([]int{1, 5, 2, 1, 4, 0}), "default example")
}
