func minCostToMoveChips(position []int) int {
    m := map[int]int{}
    for _, p := range position {
        if n, ok := m[p]; ok {
            m[p] = n+1
        } else {
            m[p] = 1
        }
    }
    ans := 65536
    for basis, _ := range m {
        buf := 0
        for k, v := range m {
            if k > basis {
                buf += v * ((k-basis) % 2)
            } else {
                buf += v * ((basis-k) % 2)
            }
        }
        if buf < ans {
            ans = buf
        }
    }
    return ans
}
