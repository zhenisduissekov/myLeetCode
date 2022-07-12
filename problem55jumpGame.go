func canJump(nums []int) bool {
    gate := len(nums)
    for i:=len(nums)-1;i>-1;i-- {
        if i + nums[i] >= gate {
            gate = i
        }
    }
    if gate == 0 {
        return true
    }
    return false
}
