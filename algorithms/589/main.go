func preorder(root *Node) []int {
    stack := []*Node{root}
    ans := []int{}
    for {
        if len(stack) == 0 {
            break
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil {
            ans = append(ans, node.Val)
            for i := len(node.Children)-1 ; i >= 0 ; i-- {
                stack = append(stack, node.Children[i])
            }
        }
    }
    return ans
}
