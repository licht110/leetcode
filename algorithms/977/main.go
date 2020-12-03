func sortedSquares(nums []int) []int {
    stack := []int{}
    for i := 0 ; i < len(nums) ; i++ {
        n := nums[i]
        if n < 0 {
            stack = append(stack, -n)
            continue
        }
        if len(stack) == 0 {
            nums[i] = square(nums[i])
            continue
        }
        if stack[len(stack)-1] > n {
            nums[i-len(stack)] = square(n)
            continue
        }
        nums[i-len(stack)] = square(stack[len(stack)-1])
        stack = stack[:len(stack)-1]
        i--
    }
    for i := 0 ; i < len(stack) ; i++ {
        nums[len(nums)-1-i] = square(stack[i])
    }
    return nums
}

func square(n int) int {
    return int(math.Pow(float64(n), 2.0))
}
