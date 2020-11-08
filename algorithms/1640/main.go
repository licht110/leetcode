object Solution {
    def canFormArray(arr: Array[Int], pieces: Array[Array[Int]]): Boolean = {
        val hash = pieces.map(x => (x(0) -> x)).toMap
        for ( i <- 0 to arr.length-1 ) {
            val a = hash.get(arr(i))
            if (a == None) {
                return false
            }
            val l = a.fold(Array(-1))(x => x).length
            if (arr.toSeq.slice(i, i+l) sameElements a) {
                i = i + l - 1
            } else {
                return false
            }
        }
        true
    }
}
