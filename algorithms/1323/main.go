func maximum69Number (num int) int {
    nine := []rune("9")[0]
    six  := []rune("6")[0]
    for i, c := range fmt.Sprintf("%v", num) {
        if c == six {
            buf := []rune(fmt.Sprintf("%v", num))
            buf[i] = nine
            ans, _ := strconv.Atoi(string(buf))
            return ans
        }
    }
    return num
}
