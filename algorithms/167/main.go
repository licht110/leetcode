func twoSum(numbers []int, target int) []int {
    m := map[int]int{}
    for i, n := range numbers {
        diff := target-n
        if index, ok := m[diff] ; ok {
            return []int{index, i+1}
        }
        m[n] = i+1
    }
    return []int{-1, -1}
}
