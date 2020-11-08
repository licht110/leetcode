object Solution {
    def decompressRLElist(nums: Array[Int]): Array[Int] = {
        (for (i <- Range(0, nums.length, 2)) yield {
            val (f, v) = (nums(i), nums(i+1))
            (for (j <- 0 to f-1) yield v)
            }.toArray
       ).toArray.flatten
    }
}
