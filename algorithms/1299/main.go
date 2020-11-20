func replaceElements(arr []int) []int {
    max := -1
    rev := []int{}
    for i := len(arr)-1 ; i >= 0 ; i-- {
        rev = append(rev, max)
        n := arr[i]
        if n > max {
            max = n
        }
    }
    ans := []int{}
    l := len(rev)
    for i := 0 ; i < len(rev) ; i++ {
        ans = append(ans, rev[l-i-1])
    }
    return ans
}
