object Solution {
    def createTargetArray(nums: Array[Int], index: Array[Int]): Array[Int] = {
        var target: List[Int] = List()
        for ( (n, i) <- (nums, index).zipped ) {
            target = target.dropRight(target.length - i) ::: List(n) ::: target.drop(i)
        }
        target.toArray
    }
}
