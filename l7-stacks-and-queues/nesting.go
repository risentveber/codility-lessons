package stacksandqueues

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/
func NestingSolution(S string) int {
	var stackLen int
	for _, char := range S {
		if char == '(' {
			stackLen++
		} else {
			stackLen--
			if stackLen < 0 {
				return 0
			}
		}
	}

	if stackLen == 0 {
		return 1
	}

	return 0
}
