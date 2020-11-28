func arrayPairSum(nums []int) int {
    sort.Slice(nums, func(i,j int) bool {return nums[i] > nums[j]})
    ans := 0
    for i := 0 ; i < len(nums) ; i += 2 {
        ans += nums[i+1]
    }
    return ans
}
