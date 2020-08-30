package timecomplexity

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
func FrogJmpSolution(X int, Y int, D int) int {
	diff := Y - X
	result := diff / D
	if diff%D != 0 {
		result++
	}

	return result
}
