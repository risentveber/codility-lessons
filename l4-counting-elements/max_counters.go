package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/
func MaxCountersSolution(N int, A []int) []int {
	base := 0
	currentMax := 0
	overBase := make(map[int]int)
	for _, operation := range A {
		if operation == N+1 {
			base = base + currentMax
			overBase = make(map[int]int)
			currentMax = 0
		} else {
			overBase[operation]++
			if overBase[operation] > currentMax {
				currentMax = overBase[operation]
			}
		}
	}
	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[i] = base + overBase[i+1]
	}

	return result
}
