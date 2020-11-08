func processQueries(queries []int, m int) []int {
    P := make([]int, m)
    for i := 1 ; i <= m ; i++ {
        P[i-1] = i
    }
    answer := []int{}
    for _, q := range queries {
        for i, v := range P {
            if v == q {
                answer = append(answer, i)
                for j := i ; j > 0 ; j-- {
                    P[j] = P[j-1]
                }
                P[0] = v
                break
            }
        }
    }
    return answer
}
