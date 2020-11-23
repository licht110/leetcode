func repeatedNTimes(A []int) int {
    m := map[int]int{}
    for i, n := range A {
        if _, ok := m[n] ; ok {
            return n
        } else {
            m[n] = i
        }
    }
    return 0
}
