func numTeams(rating []int) int {
    count := 0
    if len(rating) < 3 {
        return count
    }
    for i := 0 ; i < len(rating) ; i++ {
        for j := i+1 ; j < len(rating) ; j++ {
            for k := j+1 ; k < len(rating) ; k++ {
                if rating[i] < rating[j] && rating[j] < rating[k] {
                    count += 1
                } else if rating[i] > rating[j] && rating[j] > rating[k] {
                    count += 1
                }
            }
        }
    }
    return count
}
