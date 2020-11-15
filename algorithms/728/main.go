func selfDividingNumbers(left int, right int) []int {
    ans := []int{}
    for n := left ; n <= right ; n++ {
        if isSelfDividingNumber(n) {
            ans = append(ans, n)
        }
    }
    return ans
}

func isSelfDividingNumber(n int) bool {
    s := strconv.Itoa(n)
    for _, char := range s {
        partOfNum, err := strconv.Atoi(string(char))
        if err != nil {
            return false
        } else if partOfNum == 0 {
            return false
        } else if n % partOfNum != 0 {
            return false
        }
    }
    return true
}
