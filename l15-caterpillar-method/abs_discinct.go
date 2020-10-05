package l15caterpillarmethod

func AbsDistinctSolution(A []int) int {
	result := make(map[int]struct{})
	for _, e := range A {
		if e > 0 {
			result[e] = struct{}{}
		} else {
			result[-e] = struct{}{}
		}
	}

	return len(result)
}
