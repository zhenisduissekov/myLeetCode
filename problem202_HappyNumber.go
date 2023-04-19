func isHappy(n int) bool {
    mem := make(map[int]struct{})
    for n >= 1{
        n = getSum(n)
        if n==1 {
            return true
        }
        if _, exists := mem[n]; exists {
            return false
        } else {
            mem[n]=struct{}{}
        }
    }
    return false
}


func getSum(n int) int {
    sum := 0
    for n > 0{
        d := n%10
        sum = sum + d*d
        n = n / 10
    }
    return sum
}
