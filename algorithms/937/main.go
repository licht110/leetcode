func reorderLogFiles(logs []string) []string {
    letterLogs := []string{}
    digitLogs := []string{}
    r := regexp.MustCompile("[a-z]")
    for _, log := range logs {
        msgs := strings.Split(log, " ")
        if r.MatchString(msgs[1]) {
            letterLogs = append(letterLogs, log)
        } else {
            digitLogs = append(digitLogs, log)
        }
    }
    letterLogs = reorder(letterLogs)
    return append(letterLogs, digitLogs...)
}

func reorder(l []string) []string {
    sorter := LogSorter{}
    for _, s := range l {
        sorter = append(sorter, s)
    }
    sort.Sort(sorter)
    ret := []string{}
    for _, s := range sorter {
        ret = append(ret, s)
    }
    return ret
}

type LogSorter []string

func (l LogSorter) Len() int {
    return len(l)
}

func (l LogSorter) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l LogSorter) Less(i, j int) bool {
    msgs1 := strings.Split(l[i], " ")
    msgs2 := strings.Split(l[j], " ")
    s1, s2 := strings.Join(msgs1[1:len(msgs1)], " "), strings.Join(msgs2[1:len(msgs2)], " ")
    if s1 == s2 { return msgs1[0] < msgs2[0] }
    return s1 < s2
}
