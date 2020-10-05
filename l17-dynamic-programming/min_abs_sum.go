package l17dynamicprogramming

func MinAbsSumSolution(A []int) int {
	max, sum := 0, 0
	count := make([]int, 101)
	for i := 0; i < len(A); i++ {
		if A[i] < 0 {
			A[i] = -A[i]
		}
		if A[i] > max {
			max = A[i]
		}
		sum += A[i]
		count[A[i]]++
	}

	dp := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		dp[i] = -1
	}

	for i := 1; i <= max; i++ {
		if count[i] > 0 {
			for j := 0; j <= sum; j++ {
				if dp[j] >= 0 {
					dp[j] = count[i]
				} else if j >= i && dp[j-i] > 0 {
					dp[j] = dp[j-i] - 1
				}
			}
		}
	}

	result := sum

	for i := 0; i < sum/2+1; i++ {
		if dp[i] < 0 {
			continue
		}
		rest := sum - 2*i
		if rest < result {
			result = rest
		}
	}

	return result
}
