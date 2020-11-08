func trap(height []int) int {
        if len(height) == 0 {
                return 0
        }
        left := 0
        leftMax := height[left]
        right := len(height) - 1
        rightMax := height[right]
        ans := 0
        for {
                if left >= right {
                        break
                }
                if leftMax > rightMax {
                        ans += rightMax - height[right]
                        right -= 1
                        if rightMax < height[right] {
                                rightMax = height[right]
                        }
                } else {
                        ans += leftMax - height[left]
                        left += 1
                        if leftMax < height[left] {
                                leftMax = height[left]
                        }
                }
        }
        return ans
}
