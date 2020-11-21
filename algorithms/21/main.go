func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    } else if l2 == nil {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    }
    diff := l1.Val - l2.Val
    switch {
        case diff < 0:
            l1.Next = mergeTwoLists(l1.Next, l2)
            return l1
        default:
            l2.Next = mergeTwoLists(l1, l2.Next)
            return l2
    }
}
