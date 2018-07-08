package solution

import "sort"

func Solution(A []int) int {
    sort.Ints(A)
    length := len(A)

    max3 := A[length-1] * A[length-2] * A[length-3]
    maxmin2 := A[length-1] * A[0] * A[1]

    if max3 > maxmin2 {
    	return max3
    } else {
    	return maxmin2
    }
}
