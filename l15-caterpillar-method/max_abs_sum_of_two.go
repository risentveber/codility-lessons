package l15caterpillarmethod

import "sort"

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func MaxAbsSumOfTwoSolution(A []int) int {
	min := abs(2 * A[0])

	if len(A) == 1 {
		return min
	}

	sort.Ints(A)

	for l, r := 0, len(A)-1; l <= r; {
		if abs(A[l]+A[r]) < min {
			min = abs(A[l] + A[r])
		}

		if min == 0 {
			return 0
		}

		if abs(A[l]) > abs(A[r]) {
			l++
		} else {
			r--
		}
	}

	return min
}
