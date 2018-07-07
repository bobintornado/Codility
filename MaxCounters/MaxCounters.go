package solution

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Solution(N int, A []int) []int {
	counters := make([]int, N, N)
	max := 0
	lastMax := 0
	
	for _, val := range A {
		if val == N + 1 {
			lastMax = max
		} else {
			counters[val-1] = Max(counters[val-1], lastMax)
			counters[val-1] = counters[val-1] + 1
			if counters[val-1] > max {
				max = counters[val-1]
			}	
		}
	}

	for index, _ := range counters {
		counters[index] = Max(counters[index], lastMax)
	}

	return counters
}
