func isAlienSorted(words []string, order string) bool {
    orderMap := map[string]int{}
    for i := 0 ; i < len(order) ; i++ {
        orderMap[order[i:i+1]] = i
    }
    prev := ""
    for _, word := range words {
        if !isOrdered(orderMap, prev, word) {
            return false
        }
        prev = word
    }
    return true
}

func isOrdered(orderMap map[string]int, word1, word2 string) bool {
    for i := 0 ; i < len(word1) ; i++ {
        order1 := orderMap[word1[i:i+1]]
        if len(word2) < i+1 {
            return false            
        }
        order2 := orderMap[word2[i:i+1]]
        if order1 == order2 {
            continue
        } else if order1 < order2 {
            return true
        } else {
            return false
        }
    }
    return true
}
