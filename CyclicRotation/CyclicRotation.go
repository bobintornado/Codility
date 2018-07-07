package solution

func Solution(A []int, K int) []int {
    length := len(A)
    // if empty array or single element array
    if length < 2 {
        return A
    }
    
    k := K % length
    
    for k > 0 {
    	A = append([]int{A[length-1]}, A[:length-1]...)
    	k = k - 1
    }

   	return A
}
