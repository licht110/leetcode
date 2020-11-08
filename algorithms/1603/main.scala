class ParkingSystem(_big: Int, _medium: Int, _small: Int) {
    val typeMax = Map(1 -> _big, 2 -> _medium, 3 -> _small)
    var typeCount = Map(1 -> 0, 2 -> 0, 3 -> 0)
    
    def addCar(carType: Int): Boolean = {
        typeCount += (carType -> (typeCount(carType)+1))
        typeCount(carType) <= typeMax(carType)
    }

}
