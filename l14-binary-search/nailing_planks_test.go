package l14binarysearch

import (
	"testing"

	"github.com/tj/assert"
)

func TestNailingPlanksSolution(t *testing.T) {
	assert.Equal(t, 4,
		NailingPlanksSolution([]int{1, 4, 5, 8}, []int{4, 5, 9, 10}, []int{4, 6, 7, 10, 2}), "example")
}
