func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    curr := head
    for {
        if l1 == nil && l2 == nil {
            break
        } else if l1 == nil {
            curr.Next = l2
            l2 = l2.Next
        } else if l2 == nil {
            curr.Next = l1
            l1 = l1.Next
        } else {
            if l1.Val < l2.Val {
                curr.Next = l1
                l1 = l1.Next
            } else {
                curr.Next = l2
                l2 = l2.Next
            }
        }
        curr = curr.Next
    }
    return head.Next
}
