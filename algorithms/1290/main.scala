import scala.math.pow

object Solution {
    def getDecimalValue(head: ListNode): Int = {
        val (ret, _) = recursive(head, 0)
        ret
    }
    
    def recursive(head: ListNode, count: Int): (Int, Int) = {
        (head.next match {
            case null => (0, count)
            case _ => recursive(head.next, count+1)
        }) match {
            case (x: Int, c: Int) => (pow(2, c-count).toInt * head.x + x, c)
        }
    }
}
