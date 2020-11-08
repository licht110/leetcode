object Solution {
    def smallerNumbersThanCurrent(nums: Array[Int]): Array[Int] = {
        nums.map(x => nums.filter(y => x > y).length)
    }
}
