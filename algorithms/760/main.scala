object Solution {
    def anagramMappings(A: Array[Int], B: Array[Int]): Array[Int] = {
        val map = B.zipWithIndex.toMap
        A.map(x => map.get(x).fold(-1)(x => x) )
    }
}
