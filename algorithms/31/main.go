func nextPermutation(nums []int)  {
    for l := len(nums)-1 ; l >= 0 ; l-- {
        base := nums[l]
        prev := base
        for i := l-1 ; i >= 0 ; i-- {
            n := nums[i]
            if n < base {
                nums[i], nums[l] = nums[l], nums[i]
                buf := make([]int, len(nums[i+1:]))
                copy(buf, nums[i+1:])
                sort.Slice(buf, func(i,j int) bool { return buf[i]<buf[j]})
                copy(nums[i+1:], buf)
                return
            }
            if n < prev {
                break
            }
            prev = n
        }
    }
    sort.Slice(nums, func(i,j int) bool {return nums[i]<nums[j]})
}
