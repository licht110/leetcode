object Solution {
    def balancedStringSplit(s: String): Int = {
        var (countL, countR) = (0, 0)
        var ret = 0
        s.foreach { c =>
            if (c.toString == "R") {
                countR += 1
            } else {
                countL += 1
            }
            if (countL == countR) {
                countL = 0
                countR = 0
                ret += 1
            }
        }
        ret
    }
}
