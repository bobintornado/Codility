package solution

func Solution(A []int) int {
    length := len(A)
    m := map[int]bool{}
    for _, val := range A {
        if val > length {
            return 0
        }
        if m[val] == true {
            return 0
        }
        m[val] = true
    }
    
    return 1
}
