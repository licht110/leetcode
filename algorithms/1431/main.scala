object Solution {
    def kidsWithCandies(candies: Array[Int], extraCandies: Int): Array[Boolean] = {
        candies.map(candies.max <= _ + extraCandies)
    }
}
