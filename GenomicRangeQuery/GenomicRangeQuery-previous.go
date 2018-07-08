package solution

func Solution(S string, P []int, Q []int) []int {
	// at this position, where is previous A C G T
	index := make([][4]int, len(S), len(S))
	a := -1 
	c := -1 
	g := -1 
	t := -1 
	
	// TC N 
	for idx, ele := range S {
		switch string(ele) {
		case "A":
			a = idx
		case "C":
			c = idx
		case "G":
			g = idx
		case "T":
			t = idx
		}

		index[idx] = [4]int{a,c,g,t}
	}

	ans := make([]int, len(P), len(P)) 

	// TC M 
	for i := len(P) - 1; i >=0; i-- {
		min := P[i]
		max := Q[i]
		// TC 4
		for idx, val := range [4]int{1,2,3,4} {
			if index[max][idx] >= min {
				ans[i] = val
				break
			}
		}
	} 

	// overall TC is O(N + 4M)
	// overall SC is O(4N)
	return ans
}
