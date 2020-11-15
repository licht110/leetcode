type Node struct {
    value rune
    prev  *Node
    next  *Node
}

func lengthOfLongestSubstring(s string) int {
    if len(s) < 1 {
        return 0
    }
    currNode := &Node{
        value: []rune(s)[0],
    }
    m := map[rune]*Node{[]rune(s)[0]: currNode}
    maxLength := len(m)
    for _, char := range s[1:] {
        node, ok := m[char]
        if ok {
            if node.next != nil {
                node.next.prev = nil
            }
            for {
                delete(m, node.value)
                if node.prev == nil {
                    break
                }
                node.next = nil
                node, node.prev = node.prev, nil
            }
        }
        newNode := &Node{
            value: char,
            prev: currNode,
            next: nil,
        }
        currNode.next = newNode
        m[char] = newNode
        currLength := len(m)
        if currLength > maxLength {
            maxLength = currLength
        }
        currNode = newNode
    }
    return maxLength
}
