object Solution {
    def xorOperation(n: Int, start: Int): Int = {
        (for (i <- 0 to n-1) yield start + 2*i).fold(0)(_ ^ _)
    }
}
