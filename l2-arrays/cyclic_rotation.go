package arrays

// https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/
func CyclicRotationSolution(A []int, K int) []int {
	length := len(A)
	result := make([]int, length)
	if length == 0 {
		return result
	}

	if K < 0 {
		K += length * (-K/length + 1)
	}
	K %= length

	for i, elem := range A {
		newIndex := (i + K) % length
		result[newIndex] = elem
	}
	return result
}
