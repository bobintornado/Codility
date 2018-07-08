package solution

func Solution(A int, B int, K int) int {
	left := A / K 
	right := B / K
	
	if A % K == 0 {
        return right - left + 1	    
	}
	
	return right - left
}
