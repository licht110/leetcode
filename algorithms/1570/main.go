type SparseVector struct {
    nums []int
}

func Constructor(nums []int) SparseVector {
    return SparseVector{
        nums: nums,
    }
}

func (this *SparseVector) dotProduct(vec SparseVector) int {
    s := 0
    for i := 0; i < len(this.nums); i++ {
        if this.nums[i] != 0 && vec.nums[i] != 0 {
            s += this.nums[i] * vec.nums[i]            
        }
    }
    return s
}
