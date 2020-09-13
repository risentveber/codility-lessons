package l12euclideanalgorithm

func isEqualSets(a, b int) bool {
	if a == b {
		return true
	}
	if a == 1 || b == 1 {
		return false
	}
	abGcd := gcd(a, b)
	if abGcd == 1 {
		return false
	}
	restA := a / abGcd
	restB := b / abGcd
	for restA != 1 {
		currentGcd := gcd(restA, abGcd)
		if currentGcd == 1 {
			return false
		}
		restA = restA / currentGcd
	}

	for restB != 1 {
		currentGcd := gcd(restB, abGcd)
		if currentGcd == 1 {
			return false
		}
		restB = restB / currentGcd
	}

	return true
}

func CommonPrimeDivisorsSolution(A []int, B []int) int {
	count := 0
	for i := range A {
		if isEqualSets(A[i], B[i]) {
			count++
		}
	}

	return count
}
