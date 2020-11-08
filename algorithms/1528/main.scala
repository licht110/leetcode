object Solution {
    def restoreString(s: String, indices: Array[Int]): String = {
        indices.zipWithIndex.sortBy(_._1).map(x => s(x._2)).mkString("")
    }
}
