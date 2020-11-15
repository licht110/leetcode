func lengthOfLongestSubstring(s string) int {
    m := map[rune]bool{}
    strings := []rune{}
    maxLength := 0
    for _, char := range s {
        _, ok := m[char]
        if ok {
            m, strings = rebuild(char, m, strings)
        }
        m[char] = true
        strings = append(strings, char)
        currLength := len(strings)
        if currLength > maxLength {
            maxLength = currLength
        }
    }
    return maxLength
}

func rebuild(targetChar rune, m map[rune]bool, strings []rune) (map[rune]bool, []rune) {
    for i, char := range strings {
        delete(m, char)
        if char == targetChar {
            strings = strings[i+1:]
            break
        }
    }
    return m, strings
}
