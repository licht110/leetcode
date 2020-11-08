func cloneTree(root *Node) *Node {
    if root == nil {
        return nil
    }
    copy := &Node{}
    copy.Val = root.Val
    for _, c := range root.Children {
        copy.Children = append(copy.Children, cloneTree(c))
    }
    return copy
}
