func rangeSumBST(root *TreeNode, L int, R int) int {
    sum := 0
    if root.Left != nil && root.Val > L {
         sum += rangeSumBST(root.Left, L, R)
    }
    if root.Right != nil && root.Val < R {
        sum += rangeSumBST(root.Right, L, R)
    }
    if root.Val >= L && root.Val <= R {
        sum += root.Val
    }
    return sum
}
