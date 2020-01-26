package sorting

import "sort"

// https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/
func MaxProductOfThreeSolution(A []int) int {
	sort.Ints(A)
	l := len(A)
	max := A[l-1] * A[l-2] * A[l-3]
	negMax := A[l-1] * A[0] * A[1]
	if max > negMax {
		return max
	}
	return negMax
}
