package iterations

const base = 2

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/
func BinaryGapSolution(N int) int {
	insideGap := false
	maxCount, count := 0, 0
	for N != 0 {
		mod := N % base
		if mod == 0 {
			if insideGap {
				count++
			}
		} else {
			insideGap = true
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
		N /= base
	}
	return maxCount
}
