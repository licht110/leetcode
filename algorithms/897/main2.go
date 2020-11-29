func increasingBST(root *TreeNode) *TreeNode {
    return recursive(root, nil)
}

func recursive(curr *TreeNode, root *TreeNode) *TreeNode {
    if curr == nil {
        return root
    }
    smallest := recursive(curr.Left, curr)
    curr.Left = nil
    curr.Right = recursive(curr.Right, root)
    return smallest
}
