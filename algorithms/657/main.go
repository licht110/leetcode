func judgeCircle(moves string) bool {
    countY, countX := 0, 0
    for i := 0 ; i < len(moves) ; i++ {
        switch moves[i:i+1] {
        case "R":
            countX += 1
        case "L":
            countX -= 1
        case "U":
            countY += 1
        case "D":
            countY -= 1
        }
    }
    if countX == 0 && countY == 0 {
        return true
    } else {
        return false
    }
}
