func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    root := &TreeNode{}
    queue := [][]*TreeNode{
        []*TreeNode{t1, t2, root},
    }
    for {
        if len(queue) == 0 {
            break
        }
        p := queue[0]
        queue = queue[1:]
        t1, t2, curr := p[0], p[1], p[2]
        var (
            t1Left  *TreeNode
            t1Right *TreeNode
            t2Left  *TreeNode
            t2Right *TreeNode
        )
        if t1 != nil {
            curr.Val += t1.Val
            t1Left, t1Right = t1.Left, t1.Right
        }
        if t2 != nil {
            curr.Val += t2.Val
            t2Left, t2Right = t2.Left, t2.Right
        }
        if t1Left != nil || t2Left != nil {
            newNode := &TreeNode{}
            curr.Left = newNode
            queue = append(queue, []*TreeNode{t1Left, t2Left, newNode})
        }
        if t1Right != nil || t2Right != nil {
            newNode := &TreeNode{}
            curr.Right = newNode
            queue = append(queue, []*TreeNode{t1Right, t2Right, newNode})
        }
    }
    return root
}
