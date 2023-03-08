func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		count := get(str)
		if _, exists := m[count]; !exists {
			m[count] = []string{str}
		} else {
			m[count] = append(m[count], str)
		}
	}
	var result [][]string
	for _, mVal := range m {
		result = append(result, mVal)
	}

	return result
}

func get(str string) [26]int {
	count := [26]int{}
	for _, val := range str {
		count[int(val)-97] += 1
	}
	return count
}
