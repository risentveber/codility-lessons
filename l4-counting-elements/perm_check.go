package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/
func PermCheckSolution(A []int) int {
	elemFound := make(map[int]bool)

	for _, elem := range A {
		elemFound[elem] = true
	}

	for i := range A {
		if !elemFound[i+1] {
			return 0
		}
	}

	return 1
}
