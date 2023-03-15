import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
    m := make(map[int]struct{})
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
        if _, exists := m[i]; exists {
            continue
        }
        m[i] = struct{}{}
		for j < k {
			switch {
			case 0-nums[i] < nums[j]+nums[k]:
				k--
			case 0-nums[i] > nums[j]+nums[k]:
				j++
			case 0-nums[i] == nums[j]+nums[k]:
				temp := []int{nums[i], nums[j], nums[k]}
				if !contains(res, temp) {
					res = append(res, temp)
				}
				j++
			}
		}
	}
	return res
}

func contains(res [][]int, temp []int) bool {
	for _, val := range res {
		if val[0] == temp[0] && val[1] == temp[1] && val[2] == temp[2] {
			return true
		}
	}
	return false
}
