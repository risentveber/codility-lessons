package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
func FrogRiverOneSolution(X int, A []int) int {
	currentPosition := 0
	ableToGo := make(map[int]bool)
	for time, position := range A {
		ableToGo[position] = true

		for ableToGo[currentPosition+1] {
			currentPosition++
		}

		if currentPosition == X {
			return time
		}
	}

	return -1
}
