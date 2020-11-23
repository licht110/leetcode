func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for {
        if i > j {
            break
        }
        sum := numbers[i]+numbers[j]
        if sum == target {
            return []int{i+1,j+1}
        } else if sum > target {
            j -= 1
        } else {
            i += 1
        }
    }
    return []int{}
}
