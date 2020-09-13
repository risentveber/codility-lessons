package l13fibonaccinumbers

func LadderSolution(A []int, B []int) []int {
	L := len(A)
	max := 1 << 30
	fib := make([]int, L+1)
	result := make([]int, L)
	if L == 0 {
		return nil
	}
	if L == 1 {
		return []int{1}
	}
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < L+1; i++ {
		fib[i] = (fib[i-1] + fib[i-2]) % max
	}

	for i := 0; i < L; i++ {
		result[i] = fib[A[i]] % (1 << uint(B[i]))
	}

	return result
}
