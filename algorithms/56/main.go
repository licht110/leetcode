func merge(intervals [][]int) [][]int {
    ans := [][]int{}
    sort.Slice(intervals, func(i,j int) bool {return intervals[i][0]<intervals[j][0]})
    buf := intervals[0]
    for i := 1 ; i < len(intervals); i++ {
        interval := intervals[i]
        start, end := interval[0], interval[1]
        if start <= buf[1] {
            if end > buf[1] {
                buf[1] = end
            }
        } else {
            ans = append(ans, buf)
            buf = interval
        }
    }
    ans = append(ans, buf)
    return ans
}
