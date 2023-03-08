func productExceptSelf(nums []int) []int {
    var res []int

    prefix := 1
    for i := 0; i<len(nums);i++ {
        res = append(res, prefix)
        prefix *= nums[i]
    }

    postfix := 1
    for i :=len(nums)-1;i>-1;i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}
