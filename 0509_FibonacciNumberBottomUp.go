
func fib(n int) int {
    m := make(map[int]int)
    m[0] = 0
    m[1] = 1
    for i:=2;i<n+1;i++ {
        m[i] = m[i-1] + m[i-2]
    }
    return m[n]
}
