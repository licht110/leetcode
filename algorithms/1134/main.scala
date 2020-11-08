import scala.math.pow

object Solution {
    def isArmstrong(N: Int): Boolean = {
        (for (n <- N.toString) yield {println(n.toString.toInt, N.toString.length);pow(n.toString.toInt, N.toString.length)}).sum.toInt == N
    }
}
