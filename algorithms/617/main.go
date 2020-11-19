func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    newNode := &TreeNode{}
    var (
        t1Left *TreeNode
        t2Left *TreeNode
        t1Right *TreeNode
        t2Right *TreeNode
    )
    if t1 != nil {
        newNode.Val += t1.Val
        t1Left, t1Right = t1.Left, t1.Right
    }
    if t2 != nil {
        newNode.Val += t2.Val
        t2Left, t2Right = t2.Left, t2.Right
    }
    newNode.Left = mergeTrees(t1Left, t2Left)
    newNode.Right = mergeTrees(t1Right, t2Right)
    return newNode
}
