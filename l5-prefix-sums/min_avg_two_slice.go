package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/
func MinAvgTwoSliceSolution(A []int) int {
	sums := make([]int64, len(A))
	for i, elem := range A {
		if i == 0 {
			sums[i] = int64(elem)
		} else {
			sums[i] = sums[i-1] + int64(elem)
		}
	}
	length := len(A)

	minCount := int64(100000)
	minSum := minCount * 10000
	minStart := 0

	for start := 0; start < length-1; start++ {
		for end := start + 1; end < length && end-start < 3; end++ {
			count := end - start + 1
			sum := sums[end]
			if start != 0 {
				sum -= sums[start-1]
			}
			// sum/count < minSum/minCount
			if sum*minCount < minSum*int64(count) {
				minSum = sum
				minCount = int64(count)
				minStart = start
			}
		}
	}

	return minStart
}
