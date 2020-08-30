package l11sieveoferatosthenes

func CountSemiprimesSolution(N int, P []int, Q []int) []int {
	sieve := make([]int, N+1)
	for i := 2; i*i <= N; i++ {
		for ; sieve[i] != 0 && i*i <= N; i++ {
		}
		for k := 2; k*i <= N; k++ {
			sieve[k*i] = i
		}
	}

	counts := make([]int, N+1) // count from 1 to i of semiprimes
	for i := 1; i <= N; i++ {
		// checks that i isn't prime and division by smallest prime divider is
		if sieve[i] != 0 && sieve[i/sieve[i]] == 0 {
			counts[i] = counts[i-1] + 1
		} else {
			counts[i] = counts[i-1]
		}
	}
	result := make([]int, len(P))

	for i := 0; i < len(P); i++ {
		result[i] = counts[Q[i]] - counts[P[i]-1]
	}

	return result
}
