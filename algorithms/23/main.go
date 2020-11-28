func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    list1 := lists[0]
    for i := 1; i < len(lists) ; i++ {
        list2 := lists[i]
        if list1 == nil {
            list1 = list2
            continue
        } else if list2 == nil {
            continue
        }
        if list1.Val > list2.Val {
            list1, list2 = list2, list1
        }
        join(list1, list2)
    }
    return list1
}

func join(list1, list2 *ListNode) {
    for {
        if list1.Next == nil {
            list1.Next = list2
            break
        }
        if list1.Next.Val > list2.Val {
            list1.Next, list2 = list2, list1.Next
        }
        list1 = list1.Next
    }
}
