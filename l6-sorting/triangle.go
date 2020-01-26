package sorting

import "sort"

// https://app.codility.com/programmers/lessons/6-sorting/triangle/
func TriangleSolution(A []int) int {
	aLen := len(A)
	if aLen < 3 {
		return 0
	}

	sort.Ints(A)
	diffs := make([]int64, aLen)
	diffs[aLen-1] = 5000000000
	for i := aLen - 2; i > 0; i-- {
		currentDiff := int64(A[i+1]) - int64(A[i])
		if currentDiff < diffs[i+1] {
			diffs[i] = currentDiff
		} else {
			diffs[i] = diffs[i+1]
		}
	}

	for i := 0; i < aLen-1; i++ {
		if int64(A[i]) > diffs[i+1] {
			return 1
		}
	}

	return 0
}
