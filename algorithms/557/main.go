func reverseWords(s string) string {
    ans := []string{}
    buf := ""
    for i := 0 ; i < len(s) ; i++ {
        char := s[i:i+1]
        if char == " " {
            ans = append(ans, buf)
            buf = ""
        } else {
            buf = char + buf
        }
    }
    ans = append(ans, buf)
    return strings.Join(ans, " ")
}
