func isValid(s string) bool {
    stack := []string{}
    targets := map[string]string{
        "{": "}",
        "(": ")",
        "[": "]",
    }
    target := ""
    for i := 0 ; i < len(s) ; i++ {
        char := s[i:i+1]
        if char == "{" || char == "(" || char == "[" {
            target, _ = targets[char]
            stack = append(stack, target)
        } else {
            if len(stack) == 0 {
                return false
            } else {
                target = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if char != target {
                    return false
                }
            }
        }
    }
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
