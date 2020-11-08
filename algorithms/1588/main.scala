object Solution {
    def sumOddLengthSubarrays(arr: Array[Int]): Int = {
        (for (r <- Range(1, arr.length+1, 2) ; i <- (0 to arr.length-r) )
                   yield (
                       for (j <- i to i+r-1) yield arr(j)
                  ).sum).sum
    }
}
