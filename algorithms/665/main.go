func checkPossibility(nums []int) bool {
    count := 0
    n := len(nums)
    for i := 0; i < n-1 ; i++ {
        if nums[i] > nums[i+1] {
            count += 1
            if i >= 1 && i < n-2 {
                if nums[i] > nums[i+2] && nums[i-1] > nums[i+1] {
                    count += 1
                }
            }
        }
        if count > 1 {
            return false
        }
    }
    return true
}
