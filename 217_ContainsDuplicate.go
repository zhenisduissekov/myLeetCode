// 1st way
func containsDuplicate(nums []int) bool {
    for i := range(nums) {
        for j := i + 1; j<len(nums);j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
    }
    return false
}

// 2nd way with sorting
import "sort"

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i:=0;i<len(nums)-1;i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }
    
    return false
}

// 3rd way with hash
func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, val := range(nums) {
        if _, exists := m[val]; exists {
            return true
        }
        m[val] = struct{}{}
    }
    return false
}
