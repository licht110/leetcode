func postorder(root *Node) []int {
    stack := []*Node{root}
    ans := []int{}
    for {
        if len(stack) == 0 {
            break
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil {
            for _, c := range node.Children {
                stack = append(stack, c)
            }
            ans = append(ans, node.Val)
        }
    }
    for i := 0; i < len(ans)/2 ; i++ {
        ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
    }
    return ans
}
