package l16greedyalgorithms

func TiedRopesSolution(K int, A []int) int {
	result := 0
	current := 0
	for _, l := range A {
		if current >= K {
			result++
			current = l
		} else {
			current += l
		}
	}
	if current >= K {
		result++
	}

	return result
}
