object Solution {
    def shuffle(nums: Array[Int], n: Int): Array[Int] = {
        val n = nums.length / 2
        return (for ( i <- 0 to n - 1 ) yield Array(nums(i), nums(i+n))).flatMap(x => x).toArray
    }
}
