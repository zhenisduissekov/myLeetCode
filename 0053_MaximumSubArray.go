func maxSubArray(nums []int) int {
    maxSub := nums[0]
    curSum := 0
    for i:=0;i<len(nums);i++ {
        if curSum < 0 {
            curSum = 0
        }
        curSum += nums[i]
        maxSub = max(maxSub, curSum)
    }
    return maxSub
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
