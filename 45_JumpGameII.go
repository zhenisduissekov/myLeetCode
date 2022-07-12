func jump(nums []int) int {
    l, r := 0, 0
    res := 0
    for r < len(nums)-1 {
        f := 0
        for i:=l;i<=r;i++ {
            if f < i+nums[i]+r {
                f = i + nums[i] + r
            }
        }
        l = r + 1
        r = f
        res += 1
    }
    return res
}
