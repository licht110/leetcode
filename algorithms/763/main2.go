func partitionLabels(S string) []int {
    min, max := 0, 0
    m := map[string]int{}
    ans := []int{}
    for {
        if min >= len(S) {
            break
        }
        target := S[min:min+1]
        if index, existTarget := m[target] ; existTarget {
            max = index
        } else {
            for i := len(S)-1 ; i >= min ; i-- {
                s := S[i:i+1]
                if s == target {
                    max = i
                    break
                }
                if _, ok := m[s] ; !ok {
                    m[s] = i
                }
            }
        }
        for i := min+1 ; i < max ; i++ {
            if index, ok := m[S[i:i+1]] ; ok {
                if index > max {
                    max = index
                }
            }
        }
        ans = append(ans, max-min+1)
        min = max+1
    }
    return ans
}
