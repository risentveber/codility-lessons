package leader

// https://app.codility.com/programmers/lessons/8-leader/dominator/
//Naive way - also working solution
//func Solution(A []int) int {
//	indexByVal :=  map[int]int{}
//	countByVal :=  map[int]int{}
//	for i, elem:= range A {
//		indexByVal[elem] = i
//		countByVal[elem]++
//	}
//	half := len(A)/2
//	for elem, count := range countByVal {
//		if count > half {
//			return indexByVal[elem]
//		}
//	}
//
//	return -1
//}

func DominatorSolution(A []int) int {
	stackSize, stackValue, stackIndex := 0, 0, -1
	for i, elem := range A {
		if stackSize == 0 {
			stackSize = 1
			stackValue = elem
			stackIndex = i
		} else {
			if elem == stackValue {
				stackSize++
			} else {
				stackSize--
			}
		}
	}

	if stackSize == 0 {
		return -1
	}

	total := 0
	for _, elem := range A {
		if elem == stackValue {
			total++
		}
	}

	if total > len(A)/2 {
		return stackIndex
	}

	return -1
}
