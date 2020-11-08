func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    n1, n2 := l1, l2
    root := &ListNode{}
    node := root
    carryOver := 0
    for {
        val1, val2 := 0, 0
        if n1 != nil {
            val1 = n1.Val
        }
        if n2 != nil {
            val2 = n2.Val
        }
        val := val1 + val2 + carryOver
        if val >= 10 {
            carryOver = 1
            val -= 10
        } else {
            carryOver = 0
        }
        if n1 != nil {
            n1 = n1.Next
        }
        if n2 != nil {
            n2 = n2.Next
        }
        next := &ListNode{}
        node.Val = val
        fmt.Println(node.Val)
        if n1 == nil && n2 == nil {
            if carryOver != 0 {
                next.Val = carryOver
                node.Next = next
            }
            break
        }
        node.Next = next
        node = next
    }
    return root
}
