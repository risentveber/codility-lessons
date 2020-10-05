package l15caterpillarmethod

import "sort"

func CountTrianglesSolution(A []int) int {
	sort.Ints(A)
	if len(A) < 3 {
		return 0
	}
	result := 0

	for biggest := 2; biggest < len(A); biggest++ {
		smallest := 0
		for mid := biggest - 1; smallest < mid; mid-- {
			for ; smallest < mid && A[smallest]+A[mid] <= A[biggest]; smallest++ {
			}
			if smallest < mid {
				result += mid - smallest
			}
		}
	}

	return result
}
