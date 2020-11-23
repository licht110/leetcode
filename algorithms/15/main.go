func threeSum(nums []int) [][]int {
    m := map[int][]int{}
    ans := map[int]map[int]int{}
    for i, n := range nums {
        if indices, ok := m[n] ; ok {
            m[n] = append(indices, i)
        } else {
            m[n] = []int{i}
        }
    }
    for i := 0 ; i < len(nums)-1 ; i++ {
        for j := i+1 ; j < len(nums) ; j++ {
            target := -(nums[i]+nums[j])
            if indices, ok := m[target] ; ok {
                if !duplicate(indices, i, j) {
                    ans = add(ans, []int{nums[i], nums[j], target})
                }
            }
        }
    }
    ret := [][]int{}
    for k1, v2 := range ans {
        for k2, v2 := range v2 {
            ret = append(ret, []int{k1, k2, v2})
        }
    }
    return ret
}

func duplicate(indices []int, i,j int) bool {
    for _, index := range indices {
        if index != i && index != j {
            return false
        }
    }
    return true
}

func add(ans map[int]map[int]int, set []int) map[int]map[int]int {
    sort.Slice(set, func(i,j int) bool {return set[i] < set[j]})
    i, j, k := set[0], set[1], set[2]
    if m1, ok := ans[i] ; ok {
        if _, ok := m1[j] ; !ok {
            ans[i][j] = k
        }
    } else {
        ans[i] = map[int]int{
            j: k,
        }
    }
    return ans
}
