package l13fibonaccinumbers

func FibFrogSolution(A []int) int {
	max := len(A) + 1
	if max == 1 { // A is empty
		return 1
	}
	fib := []int{1, 1}
	for i := 2; ; i++ {
		next := fib[i-2] + fib[i-1]
		if next > max {
			break
		}
		fib = append(fib, next)
	}
	A = append([]int{1}, A...)
	A = append(A, 1)
	size := len(A)
	result := make([]int, size)
	for i := 1; i < size; i++ {
		result[i] = -1
		if A[i] == 0 { // current has no leaf
			continue
		}

		for _, jump := range fib {
			prev := i - jump
			if prev < 0 {
				break
			}

			if result[prev] == -1 {
				continue
			}

			if result[i] == -1 || result[i] > result[prev]+1 {
				result[i] = result[prev] + 1
			}
		}
	}

	return result[size-1]
}
