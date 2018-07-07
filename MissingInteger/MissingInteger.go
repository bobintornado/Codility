package solution

func Solution(A []int) int {
    ans := 1
    
    record := map[int]bool{}
    for _, val := range A {
    	if val > 0 {
    		record[val] = true
    		
    		if val == ans {
    			ans = ans + 1

    			for true {
    				if _, ok := record[ans]; ok {
    					ans = ans + 1
    				} else {
    					break
    				}
    			}
    		}
    	}
    }

    return ans
}
