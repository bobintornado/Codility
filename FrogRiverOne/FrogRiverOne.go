package solution

func Solution(X int, A []int) int {
	leafMap := map[int]bool{}

	length := len(A)
	if length < X {
		return -1
	}

	for index, value := range A {
		// new leaf
		if _, ok := leafMap[value]; !ok {
		    leafMap[value] = true
			
			X = X - 1
			if X == 0 {
				return index
			}
		}
	}

	return -1
}
