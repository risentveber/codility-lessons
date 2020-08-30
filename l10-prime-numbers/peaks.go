package l10primenumbers

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func canBeSplit(blockSize, blockCount int, peakIndexes []int) bool {
	i := 0
	for _, peakIndex := range peakIndexes {
		if peakIndex/blockSize == i {
			i++
		}
		if peakIndex/blockSize < i {
			continue
		}

		return false
	}

	return i >= blockCount
}

func PeaksSolution(A []int) int {
	N := len(A)
	peakIndexes := make([]int, 0)
	for i := 1; i < N-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peakIndexes = append(peakIndexes, i)
		}
	}

	if len(peakIndexes) < 2 {
		return len(peakIndexes)
	}

	maxDistance := 1
	for i := 1; i < len(peakIndexes); i++ {
		maxDistance = max(peakIndexes[i]-peakIndexes[i-1], maxDistance)
	}

	blockSize := max(peakIndexes[0]+1, N-peakIndexes[len(peakIndexes)-1]) // not less that left and right parts without peaks
	blockSize = max(blockSize, maxDistance/2)                             // not less then the half of max distance
	blockSize = max(blockSize, N/len(peakIndexes))                        // not less then average

	for ; blockSize <= N; blockSize++ {
		if N%blockSize != 0 {
			continue
		}
		blocksCount := N / blockSize
		if canBeSplit(blockSize, blocksCount, peakIndexes) {
			return blocksCount
		}
	}

	return 1
}
