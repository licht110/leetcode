func heightChecker(heights []int) int {
    ordered := make([]int, len(heights))
    copy(ordered, heights)
    sort.Slice(ordered, func(i,j int) bool {return ordered[i]<ordered[j]})
    count := 0
    for i := 0 ; i < len(heights) ; i++ {
        if heights[i] != ordered[i] {
            count++
        }
    }
    return count
}
