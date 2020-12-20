func subarraySum(nums []int, k int) int {
    m := map[int]int{}
    sum := 0
    ans := 0
    for _, n := range nums {
        sum += n
        if sum == k {
            ans++
        }
        diff := sum - k
        if count, ok := m[diff]; ok {
            ans += count
        }
        if _, ok := m[sum]; ok {
            m[sum] += 1
        } else {
            m[sum] = 1
        }
    }
    return ans
}
