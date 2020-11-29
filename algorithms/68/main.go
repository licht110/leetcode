func fullJustify(words []string, maxWidth int) []string {
    ans := []string{}
    buf := []string{}
    for {
        if len(words) == 0 {
            break
        }
        word := words[0]
        words = words[1:]
        buf = append(buf, word)
        if len(strings.Join(buf, " ")) > maxWidth {
            ans = append(ans, fullfillSpaces(buf[:len(buf)-1], maxWidth))
            buf = buf[len(buf)-1:]
        }
    }
    lastLine := fullfillLastLine(strings.Join(buf, " "), maxWidth)
    ans = append(ans, lastLine)
    return ans
}

func fullfillSpaces(wordsInLine []string, maxWidth int) string {
    index := 0
    maxIndex := len(wordsInLine)-1
    for {
        if len(strings.Join(wordsInLine, "")) >= maxWidth {
            break
        }
        wordsInLine[index] += " "
        index += 1
        if index >= maxIndex {
            index = 0
        }
    }
    return strings.Join(wordsInLine, "")
}

func fullfillLastLine(lastLine string, maxWidth int) string {
    for {
        if len(lastLine) >= maxWidth {
            break
        }
        lastLine += " "
    }
    return lastLine
}
