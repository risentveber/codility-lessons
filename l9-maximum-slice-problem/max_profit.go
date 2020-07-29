package maxsliceproblem

func MaxProfitSolution(A []int) int {
	result := 0
	maxEndingInI := 0
	for i := 1; i < len(A); i++ {
		delta := A[i] - A[i-1]
		maxEndingInI = max(0, maxEndingInI+delta)
		result = max(maxEndingInI, result)
	}

	return result
}
