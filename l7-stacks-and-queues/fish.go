package stacksandqueues

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/fish/
func FishSolution(A []int, B []int) int {
	stack := make([]int, 0)
	result := 0
	for i := range A {
		size, direction := A[i], B[i]
		if direction == 1 { // upstream
			stack = append(stack, size)
		} else {
			for len(stack) != 0 {
				top := stack[len(stack)-1]
				if top > size {
					break
				}
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				result++
			}
		}
	}
	return result + len(stack)
}
