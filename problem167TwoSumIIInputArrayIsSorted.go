func twoSum(numbers []int, target int) []int {
    m := map[int]int{}
    for j:=1;j<len(numbers);j++ {
        if target - numbers[0] < numbers[j] {
            break
        }
        if target - numbers[0] == numbers[j] {
            return []int{1, j+1}
        }

        m[target-numbers[j]] = j
    }

    for k:=1;k<len(numbers)-1;k++ {
        if val, exists := m[numbers[k]]; exists {
            return []int{k+1,val+1}
        }
    }
    return []int{}
}
