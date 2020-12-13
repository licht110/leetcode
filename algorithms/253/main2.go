type Interval struct {
    start int
    end   int
}

type PriorityQueue []*Interval

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].end < pq[j].end
}

func (pq PriorityQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Interval))
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
    pq := PriorityQueue{
        &Interval{
            start: intervals[0][0],
            end: intervals[0][1],
        },
    }
    heap.Init(&pq)
    for i := 1 ; i < len(intervals) ; i++ {
        start, end := intervals[i][0], intervals[i][1]
        head := heap.Pop(&pq).(*Interval)
        if start < head.end {
            heap.Push(&pq, head)
        }
        heap.Push(&pq, &Interval{
            start: start,
            end: end,
        })
    }
    return pq.Len()
}
