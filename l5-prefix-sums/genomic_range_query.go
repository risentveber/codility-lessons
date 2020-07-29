package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/
func GenomicRangeQuerySolution(S string, P []int, Q []int) []int {
	keys := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	counters := make(map[rune][]int, 4)
	length := len(S)
	for key := range keys {
		counters[key] = make([]int, length)
	}

	for index, char := range S {
		for key := range keys {
			keyCounter := counters[key]
			if key == char { // nolint:nestif
				if index == 0 {
					keyCounter[index] = 1
				} else {
					keyCounter[index] = keyCounter[index-1] + 1
				}
			} else { // the same
				if index == 0 {
					keyCounter[index] = 0
				} else {
					keyCounter[index] = keyCounter[index-1]
				}
			}
		}
	}

	results := make([]int, len(P))

	for queryIndex := range P {
		l := P[queryIndex]
		r := Q[queryIndex]
		min := 5

		for key, val := range keys {
			keyCounter := counters[key]
			count := 0
			if l != 0 {
				count -= keyCounter[l-1]
			}
			count += keyCounter[r]
			if count != 0 && val < min {
				min = val
			}
		}

		results[queryIndex] = min
	}

	return results
}
