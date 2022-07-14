func containsNearbyDuplicate(nums []int, k int) bool {
    for i:=0;i<len(nums);i++ {
        for j:=i+1;j<len(nums);j++ {
            if nums[i] == nums[j] && j-i <= k {
                return true
            }
        }
    }
    return false
}
