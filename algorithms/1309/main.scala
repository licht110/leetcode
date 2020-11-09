object Solution {
    def freqAlphabets(s: String): String = {
        val alphabetMap = (a to z).zipWithIndex.map(x => (x._2+1, x._1.toString)).toMap
        recursive(s.split("").toList, alphabetMap)
    }
    def recursive(s: List[String], m: Map[Int,String]): String = {
        s match {
            case string1 :: string2 :: "#" :: tail => m.getOrElse((string1 + string2).toInt, "") + recursive(tail, m)
            case string :: tail => m.getOrElse(string.toInt, "") + recursive(tail, m)
            case _ => ""
        }
    }
}
