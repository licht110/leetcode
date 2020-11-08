func maxSubArray(nums []int) int {
    max := nums[0]
    curr := 0
    for _, n := range nums {
        curr += n
        if curr > max {
            max = curr
        }
        if curr < 0 {
            curr = 0
        }
    }
    return max
}
