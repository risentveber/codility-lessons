package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/
func PassingCarsSolution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	count := make([]int, len(A))

	for i := len(A) - 2; i >= 0; i-- {
		if A[i+1] == 1 {
			count[i] = count[i+1] + 1
		} else {
			count[i] = count[i+1]
		}
	}
	// fmt.Println(count)
	total := int64(0)
	for i, elem := range A {
		if elem == 0 {
			total += int64(count[i])
		}
	}

	if total > 1000000000 {
		return -1
	}

	return int(total)
}
