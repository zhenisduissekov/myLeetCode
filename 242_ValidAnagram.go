func isAnagram(s string, t string) bool {
    if sortStrings(s) == sortStrings(t) {
        return true
    }
    return false
}

func sortStrings(s string) string {
    s1 := strings.Split(s, "")
    sort.Strings(s1)
    return strings.Join(s1, "")
}
