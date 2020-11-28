func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    profit := 0
    min, max := prices[0], prices[0]
    for i := 1 ; i < len(prices) ; i++ {
        p := prices[i]
        if p < min {
            min, max = p, p
        } 
        if p > max {
            max = p
            if max - min > profit {
                profit = max - min
            }
        }
    }
    return profit
}
