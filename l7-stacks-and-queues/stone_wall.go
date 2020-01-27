package stacksandqueues

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/
func StoneWallSolution(H []int) int {
	stack := make([]int, 0)
	count := 0

	for _, h := range H {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if top > h {
				count++
				stack = stack[:len(stack)-1]
			} else if top < h {
				stack = append(stack, h)
			} else {
				break
			}
		}

		if len(stack) == 0 {
			stack = append(stack, h)
		}
	}

	return count + len(stack)
}
