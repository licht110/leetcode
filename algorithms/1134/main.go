func isArmstrong(N int) bool {
    s := fmt.Sprintf("%v", N)
    l := float64(len(s))
    sum := 0
    for _, c := range s {
        p, _ := strconv.Atoi(string(c))
        sum += int(math.Pow(float64(p), l))
    }
    return N == sum
}
