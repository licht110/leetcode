func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }
    result := ""
    r := reverse(s)
    for i := 0 ; i < len(s) ; i++ {
        for j := 0 ; j < len(s) ; j++ {
            for l := len(result) ; i+l <= len(s) && j+l <= len(s) ; l++ {
                if s[i:i+l] != r[j:j+l] {
                    break
                }
                buf := s[i:i+l]
                if len(buf) > len(result) && isPalindromic(buf) {
                    result = buf
                }
            }
        }
    }
    return result
}

func reverse(s string) string {
    runes := []rune(s)
    l := len(s)/2
    for i := 0 ; i < l ; i++ {
        runes[i], runes[len(s)-i-1] = runes[len(s)-i-1], runes[i]
    }
    return string(runes)
}

func isPalindromic(s string) bool {
    if len(s) == 1 {
        return true
    }
    l := len(s)/2
    var left, right string
    if len(s) % 2 == 0 {
        left = s[0:l]
        right = s[l:len(s)]
    } else {
        left = s[0:l]
        right = s[l+1:len(s)]
    }
    switch left {
    case reverse(right):
        return true
    default:
        return false
    }
}
