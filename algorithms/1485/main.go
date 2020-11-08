func copyRandomBinaryTree(root *Node) *NodeCopy {
    hash := make(map[*Node]*NodeCopy)
    rootCopy := recursiveCopy(root, hash)
    copyRandom(root, rootCopy, hash)
    return rootCopy
}

func recursiveCopy(root *Node, hash map[*Node]*NodeCopy) *NodeCopy {
    if root == nil {
        return nil
    }
    nodeCopy := &NodeCopy{
        Val: root.Val,
    }
    hash[root] = nodeCopy
    nodeCopy.Left = recursiveCopy(root.Left, hash)
    nodeCopy.Right = recursiveCopy(root.Right, hash)
    return nodeCopy
}

func copyRandom(root *Node, rootCopy *NodeCopy, hash map[*Node]*NodeCopy) {
    for root == nil {
        return
    }
    if root.Random != nil {
        rootCopy.Random = hash[root.Random]
    }
    copyRandom(root.Right, rootCopy.Right, hash)
    copyRandom(root.Left, rootCopy.Left, hash)
    return
}
