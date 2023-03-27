func generate(numRows int) [][]int {
    res := [][]int{}
    for i:=0;i<numRows;i++ {
        if i == 0 {
            res = append(res, []int{1})
            continue
        }
        temp := []int{1}
        for j:=0;j<i-1;j++ {
            temp = append(temp, res[i-1][j]+res[i-1][j+1])
        }
        temp = append(temp, 1)
        res = append(res, temp)
    }
    return res
}
