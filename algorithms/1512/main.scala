object Solution {
    def numIdenticalPairs(nums: Array[Int]): Int = {
        nums.groupBy(x => x).values.map(x => x.length * (x.length - 1)/2).sum
    }
}
