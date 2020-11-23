func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i,j int) bool {return nums[i] < nums[j]})
    m := map[int]int{}
    buf := map[int]map[int]int{}
    for i := 0 ; i < len(nums)-1 ; i++ {
        for j := i+1 ; j < len(nums) ; j++ {
            target := -(nums[i]+nums[j])
            if _, ok := m[target]; ok {
                buf = add(buf, []int{nums[i], nums[j], target})
            }
        }
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = i
        }
    }
    return convert(buf)
}

func convert(buf map[int]map[int]int) [][]int {
    ans := [][]int{}
    for n1, m1 := range buf {
        for n2, n3 := range m1 {
            ans = append(ans, []int{n1,n2,n3})
        }
    }
    return ans
}

func add(buf map[int]map[int]int, set []int) map[int]map[int]int {
    sort.Slice(set, func(i,j int) bool {return set[i] < set[j]})
    n1, n2, n3 := set[0], set[1], set[2]
    if m1, ok := buf[n1]; ok {
        if _, ok := m1[n2]; !ok {
            buf[n1][n2] = n3
        }
    } else {
        buf[n1] = map[int]int{
            n2: n3,
        }
    }
    return buf
}
