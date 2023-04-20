func spiralOrder(matrix [][]int) []int {
	
    top, left :=0,0
    bottom, right := len(matrix), len(matrix[0])
    res := []int{}

    for top < bottom && left < right {
        for i:=left;i<right;i++ {
            res = append(res, matrix[top][i])
        }
        top++

        for i:=top;i<bottom;i++ {
            res = append(res, matrix[i][right-1])
        }
        right--

        if top >= bottom || left >= right {
            break
        }

        for i:=right-1;i>=left;i-- {
            res = append(res, matrix[bottom-1][i])
        }
        bottom--

        for i:=bottom-1;i>=top;i-- {
            res = append(res, matrix[i][left])
        }
        left++
    }
    return res
}
