class Queue():
    def __init__(self, queue: list):
        self.queue = queue

    def pop(self):
        if self.queue:
            return self.queue.pop(-1)
        else:
            return None

    def push(self, tree_node: TreeNode):
        self.queue.append(tree_node)

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        queue = Queue([cloned])
        while True:
            node = queue.pop()
            if node is None:
                break
            if node.val == target.val:
                return node
            if not node.right is None:
                queue.push(node.right)
            if not node.left is None:
                queue.push(node.left)
        return None
