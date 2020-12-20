func decodeString(s string) string {
    stack := []string{}
    for i := 0 ; i < len(s) ; i++ {
        char := s[i:i+1]
        if char == "]" {
            buf := ""
            count := ""
            for {
                poped := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if poped == "[" {
                    for {
                        re := regexp.MustCompile(`[0-9]`)
                        if len(stack) == 0 || !re.MatchString(stack[len(stack)-1]) {
                            break
                        }
                        count = fmt.Sprintf("%s%s", stack[len(stack)-1], count)
                        stack = stack[:len(stack)-1]
                    }
                    break
                }
                buf = fmt.Sprintf("%s%s", poped, buf)
            }
            countInt, err := strconv.Atoi(count)
            if err != nil {
                panic(fmt.Sprintf("unexpected string: %s", count))
            }
            decoded := decode(buf, countInt)
            for j := 0 ; j < len(decoded) ; j++ {
                stack = append(stack, decoded[j:j+1])
            }
        } else {
            stack = append(stack, char)
        }
    }
    return strings.Join(stack, "")
}

func decode(s string, count int) string {
    ret := ""
    for i := 0 ; i < count ; i++ {
        ret = fmt.Sprintf("%s%s", ret, s)
    }
    return ret
}
