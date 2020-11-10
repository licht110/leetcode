func countLetters(S string) int {
    counter, ans := 1, 1
    for i := 1; i < len(S); i++ {
        if S[i-1] == S[i] {
            counter += 1
        } else {
            counter = 1
        }
        ans += counter
    }
    return ans
}
