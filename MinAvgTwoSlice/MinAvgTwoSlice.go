package solution

import "math"

func Solution(A []int) int {
	length := len(A)
	minPos := 0
	minAvg := math.MaxFloat64

	for idx, _ := range A {
		if idx+2 <= length -1 {
			avg := float64(A[idx] + A[idx+1] + A[idx+2])/3
			if avg < minAvg {
				minAvg = avg
				minPos = idx
			}
		}

		if idx+1 <= length -1 {
			avg := float64(A[idx] + A[idx+1])/2
			if avg < minAvg {
				minAvg = avg
				minPos = idx
			}
		}	
	}
	
	return minPos
}
