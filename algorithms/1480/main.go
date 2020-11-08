func runningSum(nums []int) []int {
    ret := []int{}
    for i, _ := range nums {
        ret = append(ret, Sum(nums[0:i+1]))
    }
    return ret
}

func Sum(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}
