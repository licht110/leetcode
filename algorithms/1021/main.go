func removeOuterParentheses(S string) string {
    count := 0
    buf := []int{}
    left := []rune("(")[0]
    for i, s := range S {
        if s == left {
            count += 1
        } else {
            count -= 1
        }
        if count == 0 {
            buf = append(buf, i)
        }
    }
    start := 1
    ans := ""
    for _, i := range buf {
        ans += S[start:i]
        start = i+2
    }
    return ans
}
