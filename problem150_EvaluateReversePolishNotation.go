import "strconv"
func evalRPN(tokens []string) int {
    var stack []int
    result := 0
    for _, token := range(tokens) {
        if strings.Contains("+-*/", token) {
            if len(stack) > 1 {
               result = doArithOper(stack[len(stack)-2], stack[len(stack)-1], token)
               stack = stack[:len(stack)-2] 
               stack = append(stack, result)
            } else {
                return result
            }
        } else {
            value, _ := strconv.Atoi(token)
            stack = append(stack, value)
        }
    }
    return stack[len(stack)-1]
}


func doArithOper(a, b int, token string) int {
    fmt.Println("do", a, token, b)
    switch {
        case token == "+":
            return a+b
        case token == "-":
            return a-b
        case token == "/":
            return a/b
        case token == "*":
            return a*b    
    }
    return a
}
