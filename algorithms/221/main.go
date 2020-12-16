func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    one := []byte("1")[0]
    maxCount := 0
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            a := matrix[i][j]
            if a == one {
                count := findArea(matrix, i, j, one)
                if count > maxCount {
                    maxCount = count
                }
            }
        }
    }
    return int(math.Pow(float64(maxCount), 2.0))
}

func findArea(matrix [][]byte, i, j int, one byte) int {
    m := len(matrix)
    n := len(matrix[0])
    count := 1
    for {
        line := i+count
        column := j+count
        if line >= m || column >= n {
            return count
        }
        for l := j ; l < column ; l++ {
            a := matrix[line][l]
            if a != one {
                return count
            }
        }
        for k := i ; k < line ; k++ {
            a := matrix[k][column]
            if a != one {
                return count
            }
        }
        if matrix[line][column] != one {
            return count
        }
        count++
    }
}
