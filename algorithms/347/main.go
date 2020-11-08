type Num struct {
        key   int
        value int
}

type NumList []Num

func (n NumList) Len() int {
        return len(n)
}

func (n NumList) Less(i, j int) bool {
        return n[i].value < n[j].value
}

func (n NumList) Swap(i, j int) {
        n[i], n[j] = n[j], n[i]
}

func topKFrequent(nums []int, k int) []int {
        tmp := map[int]int{}
        for _, i := range nums {
                if _, ok := tmp[i]; ok {
                        tmp[i] += 1
                } else {
                        tmp[i] = 1
                }
        }
        numList := NumList{}
        for key, value := range tmp {
                numList = append(numList, Num{key: key, value: value})
        }
        sort.Sort(sort.Reverse(numList))

        ret := []int{}
        for _, n := range numList[0:k] {
                ret = append(ret, n.key)
        }
        return ret
}
