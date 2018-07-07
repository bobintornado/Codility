package solution

func Solution(A []int) int {
    odd := 0
    for _, val := range A {
        odd = odd ^ val
    }
    
    return odd;
}
