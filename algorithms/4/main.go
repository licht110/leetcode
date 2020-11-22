func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) == 0 {
        return median(nums2)
    } else if len(nums2) == 0 {
        return median(nums1)
    }
    conbined := []int{}
    n1 := nums1[0]
    nums1 = nums1[1:]
    n2 := nums2[0]
    nums2 = nums2[1:]
    for {
        if n1 < n2 {
            conbined = append(conbined, n1)
            if len(nums1) == 0 {
                conbined = append(conbined, append([]int{n2}, nums2...)...)
                break
            }
            n1 = nums1[0]
            nums1 = nums1[1:]
        } else {
            conbined = append(conbined, n2)
            if len(nums2) == 0 {
                conbined = append(conbined, append([]int{n1}, nums1...)...)
                break
            }
            n2 = nums2[0]
            nums2 = nums2[1:]
        }
    }
    return median(conbined)
}

func median(nums []int) float64 {
    l := len(nums)/2
    if len(nums) % 2 == 0 {
        return (float64(nums[l-1]) + float64(nums[l])) / 2
    } else {
        return float64(nums[l])
    }
}
