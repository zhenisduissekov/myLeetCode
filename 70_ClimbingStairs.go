func climbStairs(n int) int {
    one, two := 1, 0
    for i:=n;i>0;i-- {
        temp := one
        one = one +two
        two = temp
    }
    return one
}
