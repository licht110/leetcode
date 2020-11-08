func diagonalSum(mat [][]int) int {
    matLength := len(mat)
    l := matLength-1
    index := matLength/2
    sum := 0
    for i := 0 ; i < index ; i++ {
        j := l-i
        sum += mat[i][i] + mat[i][j] + mat[j][i] + mat[j][j]
    }
    if matLength % 2 == 1 {
        sum += mat[index][index]
    }
    return sum
}
