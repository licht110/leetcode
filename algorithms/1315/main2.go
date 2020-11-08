type TreeNodeWithParent struct {
    self *TreeNode
    parent *TreeNodeWithParent
}

type Stack struct {
    stack []*TreeNodeWithParent
}

func (s *Stack) Pop() *TreeNodeWithParent {
    if len(s.stack) == 0 {
        return nil
    }
    head := s.stack[0]
    if len(s.stack) == 1 {
        s.stack = s.stack[0:0]
    } else {
        s.stack = s.stack[1:len(s.stack)]    
    }
    return head
}

func (s *Stack) Push(n *TreeNodeWithParent) {
    s.stack = append([]*TreeNodeWithParent{n}, s.stack...)
}

func sumEvenGrandparent(root *TreeNode) int {
    stack := Stack{
        stack: []*TreeNodeWithParent{
            &TreeNodeWithParent{
                self: root,
                parent: nil,
            },
        },
    }
    sum := 0
    for {
        node := stack.Pop()
        if node == nil {
            break
        }
        grandParent := getGrandParent(node)
        if grandParent != nil && grandParent.self.Val % 2 == 0 {
            sum += node.self.Val
        }
        if node.self.Right != nil {
            stack.Push(
                &TreeNodeWithParent{
                    self: node.self.Right,
                    parent: node,
                },
            )
        }
        if node.self.Left != nil {
            stack.Push(
                &TreeNodeWithParent{
                    self: node.self.Left,
                    parent: node,
                },
            )
        }
    }
    return sum
}

func getGrandParent(node *TreeNodeWithParent) *TreeNodeWithParent {
    if node == nil {
        return nil
    } else if node.parent == nil {
        return nil
    }
    return node.parent.parent
}
