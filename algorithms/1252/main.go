func oddCells(n int, m int, indices [][]int) int {
    oddCells := make([][]int, n)
    for i := 0 ; i < n ; i++ {
        oddCells[i] = make([]int, m)
    }
    for _, cell := range indices {
        tmpN := cell[0]
        tmpM := cell[1]
        for i := 0 ; i < n ; i++ {
            oddCells[i][tmpM] += 1
        }
        for j := 0 ; j < m ; j++ {
            oddCells[tmpN][j] += 1
        }
    }
    counter := 0
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < m ; j++ {
            if oddCells[i][j] % 2 == 1 {
                counter += 1
            }
        }
    }
    return counter
}
