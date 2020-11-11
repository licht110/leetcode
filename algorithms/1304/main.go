func sumZero(n int) []int {
    ans := []int{}
    for i := 1 ; i < n ; i++ {
        ans = append(ans, i)
    }
    ans = append(ans, -sum(ans))
    return ans
}

func sum(ans []int) int {
    sum := 0
    for _, n := range ans {
        sum += n
    }
    return sum
}
