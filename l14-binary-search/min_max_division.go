package l14binarysearch

func MinMaxDivisionSolution(K int, _ int, A []int) int {
	max := 0
	for _, item := range A {
		if item > max {
			max = item
		}
	}
	left, right := max, (len(A)/K+1)*max // basic heuristic assumptions

	for left < right {
		current := (left + right) / 2
		prevSum := 0
		chunkCount := 1
		for _, item := range A {
			if prevSum+item <= current {
				prevSum = prevSum + item
			} else {
				prevSum = item
				chunkCount++
			}
		}

		if chunkCount <= K {
			right = current
		} else {
			left = current + 1
		}
	}

	return right
}
