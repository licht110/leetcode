func generateTheString(n int) string {
    ans := ""
    for i := 0 ; i < n-1 ; i++ {
        ans += "a"
    }
    if n % 2 == 1 {
        ans += "a"
    } else {
        ans += "b"
    }
    return ans
}
