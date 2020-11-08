func highFive(items [][]int) [][]int {
    filteredScores := map[int][]int{}
    for _, item := range items {
        id, score := item[0], item[1]
        scores, ok := filteredScores[id]
        if ok {
            scores = appendToScores(scores, score)
            if len(scores) > 5 {
                scores = scores[0:5]
            }
            filteredScores[id] = scores
        } else {
            filteredScores[id] = []int{score}
        }
    }
    ret := [][]int{}
    for id, scores := range filteredScores {
        ret = append(ret, []int{id, average(scores)})
    }
    sort.Slice(ret, func(i, j int) bool {return ret[i][0] < ret[j][0]})
    return ret
}

func appendToScores(scores []int, score int) []int {
    for i, s := range scores {
        if score > s {
            scores = append(scores[0:i+1], scores[i:]...)
            scores[i] = score
            return scores
        }
    }
    scores = append(scores, score)
    return scores
}

func average(scores []int) int {
    sum := 0
    for _, score := range scores {
        sum += score
    }
    return sum/len(scores)
}
