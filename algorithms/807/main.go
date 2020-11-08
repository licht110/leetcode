func maxIncreaseKeepingSkyline(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    skylineFromRight := make([]int, n)
    gridFromTop := makeList(m, n)
    for i, line := range grid {
        skylineFromRight[i] = max(line)
        for j, l := range line {
            gridFromTop[j][i] = l
        }
    }
    skylineFromTop := make([]int, m)
    for i, line := range gridFromTop {
        skylineFromTop[i] = max(line)
    }
    gridNew := makeList(len(grid), len(grid[0]))
    for i := 0 ; i < len(grid) ; i++ {
        for j := 0 ; j < len(grid[0]) ; j++ {
            if skylineFromTop[j] >= skylineFromRight[i] {
                gridNew[i][j] = skylineFromRight[i]
            } else {
                gridNew[i][j] = skylineFromTop[j]
            }
        }
    }
    sum_increased := sum(gridNew)
    sum_origin := sum(grid)
    return sum_increased - sum_origin
}

func sum(grid [][]int) int {
    s := 0
    for _, i := range grid {
        for _, j := range i {
            s += j
        }
    }
    return s
}

func max(nums []int) int {
    maxValue := -1
    for _, n := range nums {
        if maxValue < n {
            maxValue = n
        }
    }
    return maxValue
}

func makeList(n int, m int) [][]int {
    l := make([][]int, n)
    for i := 0 ; i < m ; i++ {
        l[i] = make([]int, m)
    }
    return l
}
