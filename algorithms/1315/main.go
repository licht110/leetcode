func sumEvenGrandparent(root *TreeNode) int {
    sum := 0
    dfs(root, nil, nil, &sum)
    return sum
}

func dfs(root *TreeNode, parent *TreeNode, grandParent *TreeNode, sum *int) {
    if grandParent !=nil && grandParent.Val % 2 == 0 {
        *sum += root.Val
    }
    if root.Right == nil && root.Left == nil {
        return
    }
    if root.Right != nil {
        dfs(root.Right, root, parent, sum)
    }
    if root.Left != nil {
        dfs(root.Left, root, parent, sum)
    }
}
