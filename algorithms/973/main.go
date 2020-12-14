type Point struct {
    coordinate []int
    distance int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    ret := old[n-1]
    *pq = old[:n-1]
    return ret
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Point))
}

func kClosest(points [][]int, K int) [][]int {
    pq := PriorityQueue{}
    heap.Init(&pq)
    for _, point := range points {
        heap.Push(&pq, &Point{
            coordinate: point,
            distance: int(math.Pow(float64(point[0]), 2.0) + math.Pow(float64(point[1]), 2.0)),
        })
    }
    ans := [][]int{}
    for i := 0 ; i < K ; i++ {
        point := heap.Pop(&pq).(*Point)
        ans = append(ans, point.coordinate)
    }
    return ans
}
