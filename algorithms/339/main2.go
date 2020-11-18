func depthSum(nestedList []*NestedInteger) int {
    sum := 0
    for _, nestedInt := range nestedList {
        sum += recursive(nestedInt, 1)
    }
    return sum
}

func recursive(nestedInt *NestedInteger, depth int) int {
    if nestedInt.IsInteger() {
        return nestedInt.GetInteger() * depth
    }
    sum := 0
    for _, n := range nestedInt.GetList() {
        sum += recursive(n, depth+1)
    }
    return sum
}
