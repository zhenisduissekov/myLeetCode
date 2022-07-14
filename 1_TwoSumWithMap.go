func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, val:= range(nums) {
        temp := target - val
        if _, exists := m[temp]; exists {
            return []int{m[temp], i}
        }
        m[val] =  i
    }
    return []int{}
}
