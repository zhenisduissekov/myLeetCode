func isPalindrome(s string) bool {

    i:=0
    j:=len(s)-1

    s = strings.ToLower(s)

    for i<j {
        r1 := rune(s[i])
        r2 := rune(s[j])
        if !isAlphanum(r1) {
            i++
            continue
        }

        if !isAlphanum(r2) {
            j--
            continue
        }

        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }

	return true
}

func isAlphanum(r rune) bool {
    if 47 < r && r < 58 {
        return true
    }

    if 96 < r && r < 123 {
        return true
    }

    return false
}
