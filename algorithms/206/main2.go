func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    for {
        if head == nil {
            break
        }
        buf := head
        head = head.Next
        buf.Next = prev
        prev = buf
    }
    return prev
}
