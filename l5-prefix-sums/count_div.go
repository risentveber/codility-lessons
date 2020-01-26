package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/
func CountDivSolution(A int, B int, K int) int {
	Amod, Acount := A%K, A/K
	Bcount := B / K
	result := Bcount - Acount
	if Amod == 0 {
		result++
	}
	return result
}
