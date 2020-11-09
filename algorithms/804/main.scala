object Solution {
    def uniqueMorseRepresentations(words: Array[String]): Int = {
        val morseCodeList = Array(".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--..")
        val m = (a to z).zipWithIndex.map( x =>
            (x._1, morseCodeList(x._2))
        ).toMap
        words.map( word =>
            word.map( c =>
                m.getOrElse(c, "")
            ).mkString("")
        ).distinct.length
    }
}
