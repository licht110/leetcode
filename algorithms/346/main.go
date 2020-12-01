type MovingAverage struct {
    queue []int
    size   int
    sum    int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{
        queue: []int{},
        size: size,
        sum: 0,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    this.sum += val
    this.queue = append(this.queue, val)
    if len(this.queue) > this.size {
        this.sum -= this.queue[0]
        this.queue = this.queue[1:]
    }
    return float64(this.sum)/float64(len(this.queue))
}
