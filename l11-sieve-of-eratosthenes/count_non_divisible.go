package l11sieveoferatosthenes

func CountNotDivisibleSolution(A []int) []int {
	N := len(A)
	distinct := make(map[int]int)
	result := make([]int, N)
	for _, elem := range A {
		distinct[elem]++
	}
	max := 2*N + 1
	sieve := make([]int, max)
	for i := 0; i < max; i++ {
		sieve[i] = N - 1
	}

	for factor, count := range distinct {
		sieve[factor] -= (count - 1)
		for i := 2; i*factor < max; i++ {
			sieve[i*factor] -= count
		}
	}

	for i := 0; i < N; i++ {
		result[i] = sieve[A[i]]
	}

	return result
}
