package l10primenumbers

func MinPerimeterRectangleSolution(N int) int {
	result := 0
	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		perimter := 2 * (i + N/i)
		if i == 1 || perimter < result {
			result = perimter
		}
	}

	return result
}
