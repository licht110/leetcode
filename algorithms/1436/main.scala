object Solution {
    def destCity(paths: List[List[String]]): String = {
        val departures = paths.map(p => p(0)).toSet
        val destinations = paths.map(p => p(1)).filterNot(departures)
        return destinations(0)
    }
}
