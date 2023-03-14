func twoSum(numbers []int, target int) []int {
    s := numbers
    m := map[int]int{}
    for j:=1;j<len(s);j++ {
        if target - s[0] < s[j] {
            break
        }
        if target - s[0] == s[j] {
            return []int{1, j+1}
        }

        m[target-s[j]] = j
    }

    for k:=1;k<len(s)-1;k++ {
        if val, exists := m[s[k]]; exists {
            return []int{k+1,val+1}
        }
    }
    return []int{}
}
