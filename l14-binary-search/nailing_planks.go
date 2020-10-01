package l14binarysearch

import "sort"

func Find(elem int, sorted []int) int {
	begin, end := 0, len(sorted)-1

	for begin < end {
		mid := (begin + end) / 2

		if sorted[mid] >= elem {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return begin
}

func NailingPlanksSolution(A []int, B []int, C []int) int {
	result := -1

	positonByNail := make(map[int]int)
	var distinct []int
	for position, nail := range C {
		_, ok := positonByNail[nail]
		if !ok {
			positonByNail[nail] = position
			distinct = append(distinct, nail)
		}
	}

	sort.Ints(distinct)

	for i := 0; i < len(A); i++ {
		cursor := Find(A[i], distinct)
		minPos := -1
		for ; cursor < len(distinct) && distinct[cursor] <= B[i]; cursor++ {
			nail := distinct[cursor]
			if nail < A[i] {
				continue
			}

			if minPos == -1 || minPos > positonByNail[nail] {
				minPos = positonByNail[nail]
			}
			if minPos < result {
				break
			}
		}
		if minPos == -1 {
			return -1
		}
		if result < minPos {
			result = minPos
		}
	}

	return result + 1
}
