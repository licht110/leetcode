func consecutiveNumbersSum(N int) int {
    ans := 1
    for n := 2 ; n < N ; n++ {
        buf := N - (n * (n-1) / 2)
        if buf <= 0 {
            break
        }
        if buf % n == 0 {
            ans += 1
        }
    }
    return ans
}
