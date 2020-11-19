func sumOfDigits(A []int) int {
    sort.Slice(A, func(i, j int) bool {return A[i] < A[j]})
    min := A[0]
    buf := min/10
    return (buf + (min - buf*10) + 1) % 2
}
