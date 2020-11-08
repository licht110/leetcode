func sumOddLengthSubarrays(arr []int) int {
    nums := generateNums(len(arr))
    sum := 0
    for i := 0 ; i < len(arr) ; i++ {
        for _, l := range nums {
            if i+l-1 <= len(arr)-1 {
                sum += Sum(arr[i:i+l])
            }
            //fmt.Println(sum, i, l)
        }
    }
    return sum
}

func generateNums(length int) []int {
    nums := []int{}
    for i := 0 ; i <= length ; i++ {
        if i % 2 == 1 {
            nums = append(nums, i)
        }
    }
    return nums
}

func Sum(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
