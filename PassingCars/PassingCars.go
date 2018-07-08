package solution

func Solution(A []int) int {
    westCount := 0
    totalPair := 0

    for i := len(A) - 1; i >= 0; i-- {
    	if A[i] == 1 {
    		westCount = westCount + 1
    	} else {
    		totalPair = totalPair + westCount
    	}

    	if totalPair > 1000000000 {
    		return -1
    	}
    }

    return totalPair
}
