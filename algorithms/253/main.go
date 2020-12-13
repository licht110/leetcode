func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Slice(intervals, func(i,j int) bool {return intervals[i][0]<intervals[j][0]})
    schedules := [][]int{intervals[0]}
    for i := 1 ; i < len(intervals) ; i++ {
        overlap := true
        interval := intervals[i]
        start := interval[0]
        for j, s := range schedules {
            if start >= s[1] {
                overlap = false
                schedules[j] = interval
                break
            }
        }
        if overlap {
            schedules = append(schedules, interval)
        }
    }
    return len(schedules)
}
