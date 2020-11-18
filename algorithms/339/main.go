func depthSum(nestedList []*NestedInteger) int {
    sum := 0
    for _, nestedInt := range nestedList {
        queue := []map[int]*NestedInteger{
            map[int]*NestedInteger{
                1: nestedInt,
            },
        }
        for {
            if len(queue) == 0 {
                break
            }
            m := queue[0]
            var depth int
            var n *NestedInteger
            for key, nested := range m {
                depth, n = key, nested
            }
            queue = queue[1:]
            if n.IsInteger() {
                sum += depth * n.GetInteger()
            } else {
                for _, i := range n.GetList() {
                    queue = append(queue, map[int]*NestedInteger{depth+1: i})
                }
            }
        }
    }
    return sum
}
