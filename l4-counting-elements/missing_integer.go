package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/
func MissingIntegerSolution(A []int) int {
	elemFound := make(map[int]bool)
	for _, elem := range A {
		elemFound[elem] = true
	}
	result := 1
	for ; elemFound[result]; result++ {
	}

	return result
}
