func productExceptSelf(nums []int) []int {
    backwards := backwardMemoization(nums)
    ans := make([]int, len(nums))
    nums = append([]int{1}, nums...)
    buf := 1
    for i := 0 ; i < len(nums)-1 ; i++ {
        buf *= nums[i]
        ans[i] = buf * backwards[i]
    }
    return ans
}

func backwardMemoization(nums []int) []int {
    backwards := make([]int, len(nums))
    backwards[len(nums)-1] = 1 
    buf := 1
    for i := len(nums)-1 ; i >= 1 ; i-- {
        buf *= nums[i]
        backwards[i-1] = buf
    }
    return backwards
}
