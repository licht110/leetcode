object Solution {
    def removeOuterParentheses(S: String): String = {
        var counter = 0
        S.zipWithIndex.map(t => {
            if (t._1.toString == "(") {
                counter += 1
            } else {
                counter -= 1
            }
            (t._2, counter)
        }).filter( t => 
            t._2 == 0
        ).fold( (1, "") )( (prev, curr) =>
            ((curr._1 + 2), prev._2 + (for (i <- prev._1 to curr._1-1) yield S(i)).mkString(""))
        )._2.toString
    }
}
