func getLonelyNodes(root *TreeNode) []int {
    ret := []int{}
    if root.Left != nil && root.Right == nil {
        ret = append(ret, getLonelyNodes(root.Left)...)
        ret = append(ret, root.Left.Val)
    } else if root.Left == nil && root.Right != nil {
        ret = append(ret, getLonelyNodes(root.Right)...)
        ret = append(ret, root.Right.Val)
    } else if root.Left != nil && root.Right != nil {
        ret = append(ret, getLonelyNodes(root.Left)...)
        ret = append(ret, getLonelyNodes(root.Right)...)
    }
    return ret
}
