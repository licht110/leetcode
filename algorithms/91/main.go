func numDecodings(s string) int {
    divided, ok := divide(s)
    if !ok {
        return 0
    }
    products := []int{}
    for _, p := range divided {
        indices := []int{}
        for i := len(p)-1 ; i >= 1 ; i-- {
            ok := canDecode(p[i-1:i+1])
            if !ok {
                indices = append(indices, i)
            }
        }
        if len(indices) == 0 || indices[len(indices)-1] != 0 {
            indices = append(indices, 0)
        }
        buf := len(p)
        nums := []int{}
        for _, n := range indices {
            nums = append(nums, fibonacci(buf - n))
            buf = n
        }
        products = append(products, product(nums))
    }
    return product(products)
}

func divide(s string) ([]string, bool) {
    buf := ""
    divided := []string{}
    for i := 0 ; i < len(s) ; i++ {
        // exit if `0` continues
        if s[i:i+1] == "0" {
            return []string{}, false
        }
        if i+2 <= len(s) && s[i+1:i+2] == "0" {
            if canDecode(s[i:i+2]) {
                divided = append(divided, buf)
                buf = ""
                i++
            } else {
                return []string{}, false
            }
        } else {
            buf += s[i:i+1]
        }
    }
    divided = append(divided, buf)
    return divided, true
}

func product(nums []int) int {
    ans := 1
    for _, n := range nums {
        ans *= n
    }
    return ans
}

func canDecode(s string) bool {
    n, ok := strconv.Atoi(s)
    if ok == nil && n >= 1 && n <= 26 {
        return true
    } else {
        return false
    }
}

func fibonacci(n int) int {
    if n <= 1 {
        return 1
    }
    a,b := 1, 1
    ans := 0
    for i := 0 ; i < n - 1 ; i++ {
        ans = a + b
        a, b = ans, a
    }
    return ans
}
