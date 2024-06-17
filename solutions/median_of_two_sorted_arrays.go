package solutions

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	merged := make([]int, totalLen)

	i1, i2 := 0, 0
	for i := 0; i < totalLen; i++ {
		if i2 >= len(nums2) || i1 < len(nums1) && nums1[i1] < nums2[i2] {
			merged[i] = nums1[i1]
			i1++
		} else {
			merged[i] = nums2[i2]
			i2++
		}
	}

	mergedMid := float64(len(merged)+1)/2.0 - 1
	midL, midR := int(math.Floor(mergedMid)), int(math.Ceil(mergedMid))

	return float64(merged[midL]+merged[midR]) / 2.0
}
