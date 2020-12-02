func canBeEqual(target []int, arr []int) bool {
    sort.Slice(target, func(i,j int) bool {return target[i]<target[j]})
    sort.Slice(arr, func(i,j int) bool {return arr[i]<arr[j]})
    return reflect.DeepEqual(target, arr)
}
