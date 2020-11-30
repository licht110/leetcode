func maximumWealth(accounts [][]int) int {
    max := 0
    for _, c := range accounts {
        buf := 0
        for _, b := range c {
            buf += b
        }
        if buf > max {
            max = buf
        }
    }
    return max
}
