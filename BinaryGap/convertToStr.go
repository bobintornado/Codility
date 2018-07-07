package solution

import "strconv"

func Solution(N int) int {
    maxGap := 0
    binaryStr := strconv.FormatInt(int64(N), 2)

    var one rune = '1'
    previousIndex := -1
    for index, char := range binaryStr {
        if char == one {
            if previousIndex == -1 {
                previousIndex = index
            } else {
                gap := index - previousIndex - 1
                if gap > maxGap {
                    maxGap = gap
                }
                previousIndex = index
            }
        }
    }

    return maxGap;
}
