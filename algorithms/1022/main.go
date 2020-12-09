func sumRootToLeaf(root *TreeNode) int {
    return recursive(root, "")
}

func recursive(root *TreeNode, s string) int {
    if root.Left != nil && root.Right != nil {
        return recursive(root.Left, s+strconv.Itoa(root.Val)) + recursive(root.Right, s+strconv.Itoa(root.Val))
    } else if root.Left != nil {
        return recursive(root.Left, s+strconv.Itoa(root.Val))
    } else if root.Right != nil {
        return recursive(root.Right, s+strconv.Itoa(root.Val))
    } else {
        return calcBit(s+strconv.Itoa(root.Val))
    }
}

func calcBit(s string) int {
    sum := 0
    for i := 0 ; i < len(s) ; i++ {
        l := len(s)-i
        if s[l-1:l] == "1" {
            sum += int(math.Pow(2.0, float64(i)))
        }
    }
    return sum
}
