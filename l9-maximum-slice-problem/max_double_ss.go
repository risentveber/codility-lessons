package maxsliceproblem

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxDoubleSliceSolution(A []int) int {
	leftMax := make([]int, len(A))
	rightMax := make([]int, len(A))
	l := len(A)

	for i := 2; i < l-1; i++ {
		leftMax[i] = max(0, leftMax[i-1]+A[i-1])
	}

	for i := l - 3; i > 0; i-- {
		rightMax[i] = max(0, rightMax[i+1]+A[i+1])
	}

	result := 0
	for i := 1; i < l-1; i++ {
		result = max(leftMax[i]+rightMax[i], result)
	}

	return result
}
