object Solution {
    def numJewelsInStones(J: String, S: String): Int = {
        (for (i <- J; j <- S) yield i == j).filter(_ == true).length
    }
}
