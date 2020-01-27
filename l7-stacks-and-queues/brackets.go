package stacksandqueues

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/
func BracketsSolution(S string) int {
	reverseMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := make([]rune, 0, 10)
	for _, char := range S {
		reverse, ok := reverseMap[char]
		if ok {
			stack = append(stack, reverse)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
