type RandomizedSet struct {
    itemMap map[int]int
    itemList []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        itemMap: map[int]int{},
        itemList: []int{},
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.itemMap[val]; ok {
        return false
    } else {
        this.itemMap[val] = len(this.itemMap)
        this.itemList = append(this.itemList, val)
        return true
    }
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if index, ok := this.itemMap[val]; ok {
        delete(this.itemMap, val)
        if len(this.itemList) <= 1 {
            this.itemList = []int{}
        } else {
            lastVal := this.itemList[len(this.itemList)-1]
            this.itemMap[lastVal] = index
            this.itemList[index] = lastVal
            this.itemList = this.itemList[:len(this.itemList)]
        }
        return true
    } else {
        return false
    }
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    randomIndex := rand.Intn(len(this.itemList))
    return this.itemList[randomIndex]
}
