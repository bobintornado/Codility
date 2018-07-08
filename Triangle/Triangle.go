package solution

import "sort"

func Solution(A []int) int {
	s := []int{}
	for _, val := range A {
		if val > 0 {
			s = append(s, val)
		}
	}

	length := len(s)
	if length < 3 {
		return 0
	}

	sort.Ints(s)

	for i := 0; i <= length - 3; i++ {
		if s[i] + s[i+1] > s[i+2] {
			return 1
		}
	}

	return 0
}
