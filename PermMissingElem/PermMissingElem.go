package solution

func Solution(A []int) int {
    // write your code in Go 1.4
    length := len(A)
    sum := (1 + length + 1) * (length  + 1) / 2 
    
    for _, val := range A {
        sum = sum - val
    }
    
    return sum
}
