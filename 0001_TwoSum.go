// I way
func twoSum(nums []int, target int) []int {
    for i:=0;i<len(nums);i++ {
        for j:=i+1;j<len(nums);j++ {
            if target == nums[i] + nums[j] {
                return []int{i, j}
            }
        }
    }
    return []int{0, 0}
}

//II way
func twoSum(nums []int, target int) []int {
	slices.Sort(nums)
	i, j := 0, len(nums)-1
	for i < j {
		switch {
		case nums[i]+nums[j] > target:
			j--
		case nums[i]+nums[j] < target:
			i++
		case nums[i]+nums[j] == target:
			return []int{i, j}
		}
	}
	return []int{}
}
