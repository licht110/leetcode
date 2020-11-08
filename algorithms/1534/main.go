func countGoodTriplets(arr []int, a int, b int, c int) int {
    ret := 0
    for i := 0 ; i < len(arr)-2 ; i++ {
        for j := i+1 ; j < len(arr)-1 ; j ++ {
            if compareAbs(arr[i]-arr[j], a) {
                for k := j+1 ; k < len(arr) ; k++ {
                    if compareAbs(arr[j]-arr[k], b) && compareAbs(arr[i]-arr[k], c) {
                        ret += 1
                    }
                }
            }
        }
    }
    return ret
}

func compareAbs(abs int, x int) bool {
    return -x <= abs && abs <= x
}
