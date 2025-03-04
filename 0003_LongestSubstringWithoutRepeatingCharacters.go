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


or


func lengthOfLongestSubstring1(s string) int {
    if len(s) < 2 {
        return len(s)
    }
    mp := map[byte]struct{}{s[0]:struct{}{}}
    i, j, mx := 0, 1, 1
    for j < len(s) {
        if _, exists := mp[s[j]]; exists {
            delete(mp, s[i])
            i++
            continue
        } 
         mp[s[j]]=struct{}{}
        j++
        mx = max(mx, len(mp))
    }
    return  mx
}


func lengthOfLongestSubstring(s string) int {
    if len(s) < 2 {
        return len(s)
    }
	i, j, mx := 0, 1, 0
	for i < len(s) && j < len(s) {
		if i != j && strings.Contains(string(s[i:j]), string(s[j])) {
				i++
		} else {
			j++
		}
		mx = max(mx, j-i)
	}
	return mx
}
