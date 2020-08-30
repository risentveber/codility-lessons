package leader

func EquiLeaderSolution(A []int) int {
	stackSize, stackValue := 0, 0
	for _, elem := range A {
		if stackSize == 0 {
			stackSize = 1
			stackValue = elem
		} else {
			if elem == stackValue {
				stackSize++
			} else {
				stackSize--
			}
		}
	}

	if stackSize == 0 {
		return 0
	}

	total := 0
	for _, elem := range A {
		if elem == stackValue {
			total++
		}
	}

	if total <= len(A)/2 {
		return 0
	}

	result := 0
	count := 0
	for index, elem := range A {
		if elem == stackValue {
			count++
		}
		if count > (index+1)/2 && total-count > (len(A)-(index+1))/2 {
			result++
		}
	}

	return result
}
