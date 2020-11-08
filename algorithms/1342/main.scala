object Solution {
    def numberOfSteps (num: Int): Int = {
        recursive(num, 0)
    }

    def recursive(num: Int, count: Int): Int = {
        num match {
            case 0 => count
            case n if n % 2 == 0 => recursive(n/2, count+1)
            case n => recursive(n-1, count+1)
        }
    }
}
