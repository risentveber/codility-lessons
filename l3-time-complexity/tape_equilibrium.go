package timecomplexity

func abs(n int64) int64 {
	if n > 0 {
		return n
	}
	return -n
}

// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
func TapeEquilibriumSolution(A []int) int {
	var totalSum int64
	for _, elem := range A {
		totalSum += int64(elem)
	}
	var sum int64
	var min int64 = -1

	for index, elem := range A {
		if index != 0 {
			tape := abs(totalSum - 2*sum)
			if tape < min || min == -1 {
				min = tape
			}
		}
		sum += int64(elem)
	}

	return int(min)
}
