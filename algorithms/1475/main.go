func finalPrices(prices []int) []int {
    ans := []int{}
    for i := 0 ; i < len(prices)-1 ; i++ {
        discount := 0
        for j := i+1 ; j < len(prices) ; j++ {
            if prices[j] <= prices[i] {
                discount = prices[j]
                break
            }
        }
        ans = append(ans, prices[i]-discount)
    }
    return append(ans, prices[len(prices)-1])
}
