func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    m1 := map[int]int{}
    m2 := map[int]int{}
    for _, n := range arr1 {
        m1[n] = 0
    }
    for _, n := range arr2 {
        m2[n] = 0
    }
    ret := []int{}
    for _, n := range arr3 {
        if _, ok1 := m1[n]; ok1 {
            if _, ok2 := m2[n]; ok2 {
                ret = append(ret, n)
            }
        }
    }
    return ret
}
