package solution

func Solution(S string, P []int, Q []int) []int {
	// at this position, what's the prefix sum (accumulation)
	index := make([][3]int, len(S), len(S))
	a := 0
	c := 0
	g := 0
	
	// TC N 
	for idx, ele := range S {
		switch string(ele) {
		case "A":
			a++
		case "C":
			c++
		case "G":
			g++
		}

		index[idx] = [3]int{a,c,g}
	}

	ans := make([]int, len(P), len(P)) 

	// TC M
	for i := len(P) - 1; i >=0; i-- {
		min := P[i]
		max := Q[i]
		
		if min == 0 {
			if index[max][0] > 0 {
				ans[i] = 1
			} else if index[max][1] > 0 {
				ans[i] = 2
			} else if index[max][2] > 0 {
				ans[i] = 3
			} else {
				ans[i] = 4
			}
		} else {
			if index[max][0] - index[min-1][0] > 0 {
				ans[i] = 1
			} else if index[max][1] - index[min-1][1] > 0 {
				ans[i] = 2
			} else if index[max][2] - index[min-1][2] > 0 {
				ans[i] = 3
			} else {
				ans[i] = 4
			}
		}
	}

	// overall TC is O(N + M)
	// overall SC is O(3N)
	return ans
}
