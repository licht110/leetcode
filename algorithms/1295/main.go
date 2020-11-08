func findNumbers(nums []int) int {
    ret := 0
    for _, n := range nums {
        if len(fmt.Sprintf("%v", n)) % 2 == 0 {
            ret += 1
        }
    }
    return ret
}
