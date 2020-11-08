func printLinkedListInReverse(head ImmutableListNode) {
    n := head
    l := []ImmutableListNode{}
    for {
        l = append(l, n)
        n = n.getNext()
        if n == nil {
            break
        }        
    }
    length_l := len(l)
    for i := length_l-1; i >= 0; i-- {
        l[i].printValue()
    }
}
