func deleteNodes(head *ListNode, m int, n int) *ListNode {
    counterM, counterN := 1, 1
    deleteFlag := false
    var (
        curr *ListNode = head
        prev *ListNode = nil
        buf *ListNode = nil
    )
    for {
        if curr == nil {
            break
        }
        if counterM > m {
            deleteFlag = true
            buf = prev
            counterM = 1
        }
        if counterN > n {
            deleteFlag = false
            buf.Next = curr
            counterN = 1
        }
        if deleteFlag {
            counterN += 1
        } else {
            counterM += 1
        }
        curr, prev = curr.Next, curr
    }
    if deleteFlag {
        buf.Next = nil
    }
    return head
}
