func hammingDistance(x int, y int) int {
    buf := x ^ y
    distance := 0
    denominator := 2147483648 // 2^31
    for {
        if buf == 0 {
            break
        }
        if  buf >= denominator {
            buf -= denominator
            distance += 1
        }
        denominator /= 2
    }
    return distance
}
