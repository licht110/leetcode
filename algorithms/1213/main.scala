object Solution {
    def arraysIntersection(arr1: Array[Int], arr2: Array[Int], arr3: Array[Int]): List[Int] = {
        arr3.filter(x => arr1.contains(x) && arr2.contains(x)).toList
    }
}
