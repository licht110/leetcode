type StackElement struct {
    node  *TreeNode
    depth int
}

type Stack struct {
    stack []StackElement
}

func (s *Stack) Pop() *StackElement{
    if len(s.stack) == 0 {
        return nil
    }
    head := s.stack[0]
    if len(s.stack) == 1 {
        s.stack = s.stack[0:0]
    } else {
        s.stack = s.stack[1:len(s.stack)]
    }
    return &head
}

func (s *Stack) Push(e StackElement) {
    s.stack = append([]StackElement{e}, s.stack...)
}

func deepestLeavesSum(root *TreeNode) int {
    stack := Stack{
        stack: []StackElement{
            StackElement{
                node: root,
                depth: 0,
            },
        },
    }
    depth, sum := 0, 0
    for {
        e := stack.Pop()
        if e == nil {
            break
        }
        if e.node.Left == nil && e.node.Right == nil {
            if depth < e.depth {
                sum = e.node.Val
                depth = e.depth
            } else if depth == e.depth {
                sum += e.node.Val
            }
        } else {
            if e.node.Right != nil {
                stack.Push(StackElement{
                    node: e.node.Right,
                    depth: e.depth + 1,
                })
            }
            if e.node.Left != nil {
                stack.Push(StackElement{
                    node: e.node.Left,
                    depth: e.depth + 1,
                })
            }
        }
    }
    return sum
}
