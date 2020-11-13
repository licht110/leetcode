func sortString(s string) string {
    ans := ""
    flag := false
    for {
        if len(s) == 0 {
            break
        }
        s = Sort(s, flag)
        buf := ""
        for i := 0 ; i < len(s) ; i++ {
            char := s[i:i+1]
            if char == buf {
                continue
            } else {
                ans += char
                s = s[0:i] + s[i+1:]
                i--
            }
            buf = char
        }
        flag = !flag
    }
    return ans
}

func Sort(s string, reverse bool) string {
    buf := []rune(s)
    if reverse {
        sort.Slice(buf, func(i, j int) bool {return buf[i] > buf[j]})
    } else {
        sort.Slice(buf, func(i, j int) bool {return buf[i] < buf[j]})
    }
    return string(buf)
}
