func increasingBST(root *TreeNode) *TreeNode {
    return recursive(root, nil)
}

func recursive(curr *TreeNode, root *TreeNode) *TreeNode {
    if curr == nil {
        return nil
    }
    var smallest *TreeNode
    if curr.Left != nil {
        smallest = recursive(curr.Left, curr)
    } else {
        smallest = curr
    }
    if curr.Right != nil {
        curr.Right = recursive(curr.Right, root)
    } else {
        curr.Right = root
    }
    curr.Left = nil
    return smallest
}
