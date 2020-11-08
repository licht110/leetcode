import "math"

func minTimeToVisitAllPoints(points [][]int) int {
    ret := 0
    p1 := points[0]
    for i := 1 ; i < len(points) ; i++ {
        p2 := points[i]
        ret += findTime(p1, p2)
        p1 = p2
    }
    return ret
}

func findTime(p1, p2 []int) int {
    x := int(math.Abs(float64(p1[0]-p2[0])))
    y := int(math.Abs(float64(p1[1]-p2[1])))
    if x == y {
        return x
    } else if x > y {
        return x
    } else {
        return y
    }
}
