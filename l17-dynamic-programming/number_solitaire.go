package l17dynamicprogramming

func NumberSolitaireSolution(A []int) int {
	l := len(A)

	for i := 1; i < l; i++ {
		max := A[i-1]
		for k := 2; k <= 6; k++ {
			if i-k >= 0 && max < A[i-k] {
				max = A[i-k]
			}
		}
		A[i] += max
	}

	return A[l-1]
}
