type RecentCounter struct {
    requests []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        requests: []int{},
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.requests = append(this.requests, t)
    threshold := t - 3000
    if threshold < 0 {
        return len(this.requests)
    }
    for i, r := range this.requests {
        if r >= threshold {
            this.requests = this.requests[i:]
            break
        }
    }
    return len(this.requests)
}
