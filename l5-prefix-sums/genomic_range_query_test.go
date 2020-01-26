package prefixsums

import (
	"testing"

	"github.com/tj/assert"
)

func TestGenomicRangeQuerySolution(t *testing.T) {
	assert.Equal(t, []int{}, GenomicRangeQuerySolution("", []int{}, []int{}), "empty")
	assert.Equal(t, []int{2, 4, 1}, GenomicRangeQuerySolution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}), "standard test")
}
