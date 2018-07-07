package solution

import "math"

func Solution(A []int) int {
    sumLeft := A[0]
    sum := 0
    for _, val := range A {
        sum = sum + val
    }
    sumRight := sum - sumLeft
    minDifference := math.Abs(float64(sumLeft - sumRight))

    p := 2
    length := len(A)
    for p < length - 1 {
    	sumLeft = sumLeft + A[p-1]
    	sumRight = sumRight - A[p-1]
    	currentDiff := math.Abs(float64(sumLeft - sumRight))
    	if currentDiff < minDifference {
    		minDifference = currentDiff
    	}

        p = p + 1
    }

    return int(minDifference)
}
