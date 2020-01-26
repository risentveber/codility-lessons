package prefixsums

import (
	"testing"

	"github.com/tj/assert"
)

func TestMinAvgTwoSliceSolution(t *testing.T) {
	testData := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		testData[i] = 10000
	}
	assert.Equal(t, 0, MinAvgTwoSliceSolution(testData), "max elem count")
}
