object Solution {
    def diagonalSum(mat: Array[Array[Int]]): Int = {
        val sum = (for (i <- 0 to mat.length/2-1) yield {
            val j = mat.length-1
            mat(i)(i) + mat(i)(j-i) + mat(j-i)(i) + mat(j-i)(j-i)
        }).sum
        if (mat.length % 2 == 1) sum + mat(mat.length/2)(mat.length/2) else sum
    }
}
