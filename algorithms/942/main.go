func diStringMatch(S string) []int {
    ans := make([]int, len(S)+1)
    max, min := len(S), 0
    var s string
    for i := 0 ; i < len(S) ; i++ {
        s = S[i:i+1]
        if s == "I" {
            ans[i] = min
            min += 1
        } else {
            ans[i] = max
            max -= 1
        }
    }
    ans[len(S)] = max
    return ans
}
