package l10primenumbers

func FlagsSolution(A []int) int {
	peakIndexes := make([]int, 0)
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peakIndexes = append(peakIndexes, i)
		}
	}

	if len(peakIndexes) < 2 {
		return len(peakIndexes)
	}

	for i := 2; true; i++ {
		rest := i
		prev := -400001 // max distance from task conditions
		for _, peakIndex := range peakIndexes {
			if peakIndex-prev >= i {
				prev = peakIndex
				rest--
			}
		}
		if rest > 0 {
			return i - 1
		}
	}

	return 0
}
