package l16greedyalgorithms

func MaxNonOverlappingSegmentsSolution(A []int, B []int) int {
	if len(A) == 0 {
		return 0
	}
	end := B[0]
	result := 1
	for i := 0; i < len(A); i++ {
		if end < A[i] {
			result++
			end = B[i]
		}
	}

	return result
}
