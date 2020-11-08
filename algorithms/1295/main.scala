object Solution {
    def findNumbers(nums: Array[Int]): Int = {
        nums.map(n => n.toString.length % 2 == 0).filter(x => x).length
    }
}
