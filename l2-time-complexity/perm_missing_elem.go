package timecomplexity

// https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
func PermMissingElem(A []int) int {
	length := int64(len(A))
	fullSum := (2 + length) * (length + 1) / 2

	sum := int64(0)
	for _, elem := range A {
		sum += int64(elem)
	}

	return int(fullSum - sum)
}
