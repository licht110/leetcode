type Logger struct {
    m map[string]int
}


/** Initialize your data structure here. */
func Constructor() Logger {
    return Logger{
        m: map[string]int{},
    }
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
        If this method returns false, the message will not be printed.
        The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    ret := false
    if t, ok := this.m[message]; ok {
        if timestamp >= t+10 {
            ret = true
            this.m[message] = timestamp
        }
    } else {
        ret = true
        this.m[message] = timestamp
    }
    return ret
}
