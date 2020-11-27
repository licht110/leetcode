func partitionLabels(S string) []int {
    queue := []string{}
    m := map[string][]int{}
    for i := 0 ; i < len(S) ; i++ {
        s := S[i:i+1]
        if l, ok := m[s] ; ok {
            m[s] = append(l, i)
        } else {
            m[s] = []int{i}
            queue = append(queue, s)
        }
    }
    min, max := 0, 0
    ans := []int{}
    for _, s := range queue {
        l, _ := m[s]
        start, end := l[0], l[len(l)-1]
        if start > max {
            ans = append(ans, max-min+1)
            min, max = start, end
        } else if end > max {
            max = end
        }
    }
    ans = append(ans, max-min+1)
    return ans
}
