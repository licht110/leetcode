type OrderedStream struct {
    stream []string
    ptr    int
}

func Constructor(n int) OrderedStream {
    return OrderedStream{
        stream: make([]string, n),
        ptr:    1,
    }
}


func (this *OrderedStream) Insert(id int, value string) []string {
    this.stream[id-1] = value
    ret := []string{}
    if id == this.ptr {
        for i := this.ptr ; i <= len(this.stream) ; i++ {
            this.ptr = i
            if this.stream[i-1] == "" {
                break
            } else {
                ret = append(ret, this.stream[i-1])
            }
        }
    }
    return ret
}
