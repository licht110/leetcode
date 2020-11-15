func countNegatives(grid [][]int) int {
    maxJ := len(grid[0])
    gridLength := len(grid)
    ans  := 0
    for i := 0 ; i < len(grid) ; i++ {
        for j := 0 ; j < maxJ ; j++ {
            if grid[i][j] < 0 {
                ans += (maxJ - j) * (gridLength - i)
                maxJ = j
                break
            }
        }
    }
    return ans
}
