func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    l:=0
    m := make(map[byte]struct{})
    m[s[0]] = struct{}{}
    max := 1
    for r:=1;r<len(s);r++ {
        for r < len(s){
            if _, ok := m[s[r]]; ok {
                delete(m, s[l])
                l += 1
                continue
            }    
            m[s[r]] = struct{}{}  
            break
        }
        if max < r-l+1 {
            max = r-l+1
        }
    }
    return max
}
