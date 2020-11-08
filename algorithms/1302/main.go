func deepestLeavesSum(root *TreeNode) int {
    _, sum :=  recursive(root, 0)
    return sum
}

func recursive(root *TreeNode, depth int) (int, int) {
    if root.Left == nil && root.Right == nil {
        return depth, root.Val
    } else if root.Right == nil {
        return recursive(root.Left, depth+1)
    } else if root.Left == nil {
        return recursive(root.Right, depth+1)
    } else {
        depthLeft, leftVal := recursive(root.Left, depth+1)
        depthRight, rightVal := recursive(root.Right, depth+1)
        if depthLeft == depthRight {
            return depthLeft, leftVal + rightVal
        } else if depthLeft > depthRight {
            return depthLeft, leftVal
        } else {
            return depthRight, rightVal
        }
    }
}
