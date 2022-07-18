func maxProfit(prices []int) int {
    L, R := 0, 1
    profit := 0
    for R < len(prices) {
        if prices[L] > prices[R] {
            L = R
        } else if profit < prices[R] - prices[L]{
            profit = prices[R] - prices[L]
        }
        R += 1
    }
    return  profit
}
