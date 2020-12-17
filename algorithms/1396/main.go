type Start struct {
    stationName string
    t int
}


type UndergroundSystem struct {
    startMap map[int]Start
    sums map[string]map[string][]int
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{
        startMap: map[int]Start{},
        sums: map[string]map[string][]int{},
    }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.startMap[id] = Start{
        stationName: stationName,
        t: t,
    }
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    start, _ := this.startMap[id]
    if m, ok := this.sums[start.stationName]; ok {
        if pair, ok2 := m[stationName]; ok2 {
            sum, count := pair[0], pair[1]
            this.sums[start.stationName][stationName] = []int{sum + (t - start.t), count+1}
        } else {
            this.sums[start.stationName][stationName] = []int{t - start.t, 1}
        }
    } else {
        this.sums[start.stationName] = map[string][]int{
            stationName: []int{t - start.t, 1},
        }
    }
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    pair := this.sums[startStation][endStation]
    sum, count := pair[0], pair[1]
    return float64(sum)/float64(count)
}
