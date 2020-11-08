func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    max, leftArray, rightArray := devide(nums)
    left := constructMaximumBinaryTree(leftArray)
    right := constructMaximumBinaryTree(rightArray)
    return &TreeNode{
        Val: max,
        Left: left,
        Right: right,
    }
}

func devide(nums []int) (int, []int, []int) {
    index := 0
    max := nums[index]
    for i, n := range nums {
        if n > max {
            index, max = i, n
        }
    }
    return max, nums[0:index], nums[index+1:len(nums)]
}
