object Solution {
    def maxSubArray(nums: Array[Int]): Int = {
        recursive(nums.toList, nums(0), 0)
    }
    
    def recursive(nums: List[Int], max: Int, curr: Int): Int = {
        nums match {
            case head :: Nil => {
                val buf = curr + head
                if (buf > max) buf else max
            }
            case head :: tail => {
                val buf = curr + head
                recursive(tail, if (buf > max) buf else max, if (buf < 0) 0 else buf)
            }
            case _ => sys.error("error")
        }
    }
}
