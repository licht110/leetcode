import scala.math._

object Solution {
    def calculateTime(keyboard: String, word: String): Int = {
        val keyMap = keyboard.zipWithIndex.map(x => (x._1, x._2)).toMap
        var sum = 0
        word.fold("A"(0))({ (c1, c2) =>
            sum += abs(keyMap.getOrElse(c1, 0) - keyMap.getOrElse(c2, 0));
            c2
        })
        sum
    }
}
