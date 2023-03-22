func isValid(s string) bool {
	var openBrackets []rune
    m := map[rune]rune{')': '(',']': '[', '}': '{'}
	for _, br := range s {
        if strings.Contains("({[", string(br)) {
            openBrackets = append(openBrackets, rune(br))
        } else if strings.Contains(")}]", string(br)) {
            if len(openBrackets) == 0 || m[rune(br)] != openBrackets[len(openBrackets)-1]{
                return false
            } else if openBrackets[len(openBrackets)-1] == m[rune(br)] {
                openBrackets = openBrackets[:len(openBrackets)-1]
            }
        }
	}
	return len(openBrackets) == 0
}
