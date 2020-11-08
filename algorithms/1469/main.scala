object Solution {
    def getLonelyNodes(root: TreeNode): List[Int] = {
        (root.left, root.right) match {
            case (left: TreeNode, null) => List(left.value) ::: getLonelyNodes(left)
            case (null, right: TreeNode) => getLonelyNodes(right) ::: List(right.value)
            case (left: TreeNode, right: TreeNode) => getLonelyNodes(left) ::: getLonelyNodes(right)
            case _ => List()
        }
    }
}
