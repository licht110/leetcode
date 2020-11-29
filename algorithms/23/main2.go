func mergeKLists(lists []*ListNode) *ListNode {
    list := []*ListNode{}
    for _, node := range lists {
        for {
            if node == nil {
                break
            }
            list = append(list, node)
            node = node.Next
        }
    }
    if len(list) == 0 {
        return nil
    }
    sort.Slice(list, func(i,j int) bool {return list[i].Val < list[j].Val})
    for i := 0 ; i < len(list)-1 ; i++ {
        node := list[i]
        node.Next = list[i+1]
    }
    list[len(list)-1].Next = nil
    return list[0]
}
