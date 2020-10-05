package l15caterpillarmethod

const MaxToCompute = 1000000000

func CountDistinctSlicesSolution(M int, A []int) int {
	positions := make([]int, M+1)
	for i := 0; i <= M; i++ {
		positions[i] = -1
	}
	result := 0

	for begin, end := 0, 0; begin < len(A); {
		for ; end < len(A); end++ {
			value := A[end]
			if positions[value] != -1 {
				break
			}
			positions[value] = end
		}

		if end == len(A) {
			// compute rest
			n := end - begin + 1
			result += n * (n - 1) / 2
			if result > MaxToCompute {
				return MaxToCompute
			}

			return result
		}

		mid := positions[A[end]]
		// left ... mid ... end(==mid)
		for i := begin; i < mid; i++ {
			positions[A[i]] = -1
		}
		positions[A[end]] = end
		result += (end - begin + end - mid) * (mid - begin + 1) / 2

		if result > MaxToCompute {
			return MaxToCompute
		}

		begin = mid + 1
		end = end + 1
	}

	return result
}
