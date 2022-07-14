
func fib(n int) int {
    m := make(map[int]int)
    m[0] = 0
    m[1] = 1
    return fibm(n, m)
}
    
func fibm(n int, m map[int]int) int {
    if val, ok := m[n]; ok {
        return val
    }
    m[n] = fibm(n-1, m) + fibm(n-2, m)
    return m[n]
}
