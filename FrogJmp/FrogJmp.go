package solution

func Solution(X int, Y int, D int) int {
    jump := (Y - X) / D
    
    if X + jump * D < Y {
        return jump + 1
    } else {
        return jump
    }
}
