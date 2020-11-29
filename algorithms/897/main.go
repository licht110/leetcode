func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    list := []*TreeNode{}
    for {
        if len(queue) == 0 {
            break
        }
        node := queue[0]
        queue = queue[1:]
        list = append(list, node)
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    sort.Slice(list, func(i,j int) bool {return list[i].Val < list[j].Val})
    for i := 0 ; i < len(list)-1 ; i++ {
        list[i].Left = nil
        list[i].Right = list[i+1]
    }
    list[len(list)-1].Left = nil
    list[len(list)-1].Right = nil
    return list[0]
}
