package maxsliceproblem

import "math"

func MaxSliceSumSolution(A []int) int {
	result := math.MinInt32
	maxEndsInI := math.MinInt32
	for i := range A {
		maxEndsInI = max(maxEndsInI+A[i], A[i])
		result = max(maxEndsInI, result)
	}

	return result
}
