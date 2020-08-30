package arrays

// https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/
func OddOccurrencesInArraySolution(A []int) int {
	odd := make(map[int]bool)
	for _, elem := range A {
		odd[elem] = !odd[elem]
	}

	for elem, isOdd := range odd {
		if isOdd {
			return elem
		}
	}

	return 0
}
