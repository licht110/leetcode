func searchBST(root *TreeNode, val int) *TreeNode {
    queue := []*TreeNode{root}
    var curr *TreeNode
    for {
        curr = queue[0]
        queue = queue[1:]
        if curr == nil || curr.Val == val {
            break
        } else if curr.Val > val {
            queue = append(queue, curr.Left)
        } else {
            queue = append(queue, curr.Right)
        }
    }
    return curr
}
