package sorting

import "sort"

type Border struct {
	x   int
	end int // 0 - start, 1 - end
}
type BorderOrder []Border

func (v BorderOrder) Len() int      { return len(v) }
func (v BorderOrder) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v BorderOrder) Less(i, j int) bool {
	return v[i].x < v[j].x || v[i].x < v[j].x && v[i].end < v[j].end
}

// https://app.codility.com/programmers/lessons/6-sorting/number_of_disc_intersections/
func NumberOfDiskIntersectionsSolution(A []int) int {
	borders := make([]Border, 0, len(A)*2)
	for x, radius := range A {
		borders = append(borders, Border{x: x - radius})
		borders = append(borders, Border{x: x + radius, end: 1})
	}

	sort.Sort(BorderOrder(borders))

	result := 0
	isOpenBefore := 0
	point := -2147483648
	openedAtPoint := 0
	closedAtPoint := 0

	for _, border := range borders {
		if point != border.x {
			result += openedAtPoint * isOpenBefore // intersections with opened before
			if openedAtPoint > 0 {
				result += (openedAtPoint - 1) * (openedAtPoint) / 2 // self intersections
			}
			if result > 10000000 {
				return -1
			}
			isOpenBefore += openedAtPoint - closedAtPoint
			point = border.x
			openedAtPoint = 0
			closedAtPoint = 0
		}

		if border.end == 1 {
			closedAtPoint++
		} else {
			openedAtPoint++
		}
	}

	return result
}
