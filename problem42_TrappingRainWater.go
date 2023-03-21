func trap(height []int) int {
	res := 0
    for i:=0;i<len(height);i++{
        maxLeft := max(height[:i])
        maxRight := max(height[i+1:])
        minBothSides := min(maxLeft, maxRight)
        if height[i] < minBothSides{
            res += minBothSides - height[i]
        }
    }
	return res
}


func max(height []int) int {
    max := 0 
    for _, val := range height {
        if max < val {
            max = val
        }
    }
    return max
}

func min (a, b int) int {
    if a > b {
        return b
    }
    return a
}
