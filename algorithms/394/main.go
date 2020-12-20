func decodeString(s string) string {
    re := regexp.MustCompile(`[0-9]+\[[a-z]+\]`)
    for {
        foundStrings := re.FindAllString(s, -1)
        if len(foundStrings) == 0 {
            break
        }
        for _, foundString := range foundStrings {
            decoded := decode(foundString)
            s = strings.Replace(s, foundString, decoded, 1)
        }
    }
    return s
}

func decode(s string) string {
    reNum := regexp.MustCompile(`^[0-9]+`)
    num, err := strconv.Atoi(reNum.FindString(s))
    if err != nil {
        panic(fmt.Sprintf("unexpected string: %s", s))
    }
    reString := regexp.MustCompile(`[a-z]+`)
    part := reString.FindString(s)
    ret := ""
    for i := 0 ; i < num ; i++ {
        ret = fmt.Sprintf("%s%s", ret, part)
    }
    return ret
}
