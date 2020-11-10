object Solution {
    def flipAndInvertImage(A: Array[Array[Int]]): Array[Array[Int]] = {
        A.map( l =>
            (for (i <- 0 to l.length-1) yield {
                l(l.length-1-i) match {
                    case 0 => 1
                    case 1 => 0
                }
            }).toArray
        )
    }
}
