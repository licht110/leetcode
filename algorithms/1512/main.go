func numIdenticalPairs(nums []int) int {
    m := map[int][]int{}
    for i, n := range nums {
        m[n] = append(m[n], i)
    }
    count := 0
    for _, v := range m {
        count += len(v) * (len(v)-1) / 2
    }
    return count
}
