package solution

import "sort"

type Events [][2]int

func (e Events) Len() int {
	return len(e)
}

func (e Events) Less(i, j int) bool {
	if e[i][0] < e[j][0] {
		return true
	}
	if e[i][0] > e[j][0] {
		return false
	}
	if e[i][1] > e[j][1] {
		return true
	}

	return false
}

func (e Events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func Solution(A []int) int {
	// pos, val
	events := Events{}

	for idx, val := range A {
		events = append(events, [2]int{idx-val, 1})
		events = append(events, [2]int{idx+val, -1})
	}

	sort.Sort(events)
	intersections, activeCircles := 0, 0

	for _, event := range events {
		if event[1] > 0 {
			intersections = intersections + activeCircles
		}
		activeCircles = activeCircles + event[1]

		if intersections > 10000000 {
			return -1
		}
	}

	return intersections
}
