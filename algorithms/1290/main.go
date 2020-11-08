import "math"

func getDecimalValue(head *ListNode) int {
    ret, _ := recursive(head, 0)
    return ret
}

func recursive(head *ListNode, count int) (int, int) {
    switch head.Next {
    case nil:
        return pow(2, 0) * head.Val, count
    default:
        ret, c := recursive(head.Next, count+1)
        return pow(2, c-count) * head.Val + ret, c
    }
}

func pow(num int, p int) int {
    return int(math.Pow(float64(num), float64(p)))
}
