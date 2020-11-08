func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i, num := range nums {
        diff := target - num
        j, ok := hash[diff]
        if ok {
            return []int{i, j}
        }
        hash[num] = i
    }
    return []int{}
}
