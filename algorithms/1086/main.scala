object Solution {
    def highFive(items: Array[Array[Int]]): Array[Array[Int]] = {
        var hash: Map[Int, List[Int]] = Map()
        items.foreach( item =>
            hash = hash.updated(item(0), (hash.getOrElse(item(0), List()) ::: List(item(1))))
        )
        hash.map( x =>
            Array(x._1, (x._2.sorted.drop(x._2.length-5).sum/5).toInt)
        )
            .toArray
            .sortBy(_(0))
    }
}
