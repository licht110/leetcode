object Solution {
    def subtractProductAndSum(n: Int): Int = {
        n.toString.map(_.toInt - 48).product - n.toString.map(_.toInt - 48).sum
    }
}
