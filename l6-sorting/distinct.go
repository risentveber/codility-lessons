package sorting

// https://app.codility.com/programmers/lessons/6-sorting/distinct/
func DistinctSolution(A []int) int {
	counter := make(map[int]bool)
	for _, elem := range A {
		counter[elem] = true
	}

	return len(counter)
}
