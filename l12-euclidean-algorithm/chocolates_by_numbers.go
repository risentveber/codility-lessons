package l12euclideanalgorithm

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}

	return b
}

func ChocolatesByNumbersSolution(N int, M int) int {
	return N / gcd(N, M)
}
