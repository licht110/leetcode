func interpret(command string) string {
    ans := ""
    for i := 0 ; i < len(command) ; i++ {
        s := command[i:i+1]
        if s == "(" {
            next := command[i+1:i+2]
            if next == ")" {
                s = "o"
                i++
            } else {
                s = "al"
                i += 3
            }
        }
        ans += s
    }
    return ans
}
