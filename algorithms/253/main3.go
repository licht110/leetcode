type PriorityQueue []int

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func(pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    ret := old[n-1]
    *pq = old[:n-1]
    return ret
}

func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Slice(intervals, func(i,j int) bool {return intervals[i][0] < intervals[j][0]})
    pq := PriorityQueue{intervals[0][1]}
    heap.Init(&pq)
    for i := 1 ; i < len(intervals) ; i++ {
        start, end := intervals[i][0], intervals[i][1]
        minEnd := heap.Pop(&pq).(int)
        if start < minEnd {
            heap.Push(&pq, minEnd)
        }

        heap.Push(&pq, end)
    }
    return pq.Len()
}
