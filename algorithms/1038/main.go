func bstToGst(root *TreeNode) *TreeNode {
    recursive(root, 0)
    return root
}

func recursive(root *TreeNode, sum int) int {
    if root == nil {
        return sum
    }
    root.Val += recursive(root.Right, sum)
    return recursive(root.Left, root.Val)
}
