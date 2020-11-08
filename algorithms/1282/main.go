func groupThePeople(groupSizes []int) [][]int {
    m := map[int]int{}
    for _, groupSize := range groupSizes {
        _, ok := m[groupSize]
        if ok {
            m[groupSize] += 1
        } else {
            m[groupSize] = 1
        }
    }
    buf := map[int][][]int{}
    for k, v := range m {
        for i := 0 ; i < v / k ; i++ {
            buf[k] = append(buf[k], []int{})
        }
    }
    for person, groupSize := range groupSizes {
        groups := buf[groupSize]
        for i := 0 ; i < len(groups) ; i++ {
            if len(groups[i]) < groupSize {
                groups[i] = append(groups[i], person)
            }
        }
    }
    ret := [][]int{}
    for _, v := range buf {
        ret = append(ret, v...)
    }
    return ret
}
