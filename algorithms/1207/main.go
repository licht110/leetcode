func uniqueOccurrences(arr []int) bool {
    m := map[int]int{}
    for _, n := range arr {
        if _, ok := m[n]; ok {
            m[n] += 1
        } else {
            m[n] = 1
        }
    }
    m2 := map[int]bool{}
    for _, v := range m {
        if _, ok := m2[v]; ok {
            return false
        } else {
            m2[v] = true
        }
    }
    return true
}
