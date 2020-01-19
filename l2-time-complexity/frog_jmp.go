package timecomplexity

func FrogJmpSolution(X int, Y int, D int) int {
	diff := Y - X
	result := diff / D
	if diff%D != 0 {
		result++
	}
	return result
}
