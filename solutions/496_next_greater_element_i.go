package solutions

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))

	var jFound, n2NextGreater int
	for i, n1 := range nums1 {
		jFound = -1
		n2NextGreater = -1
		for j, n2 := range nums2 {
			if n1 == n2 {
				jFound = j
			}

			if jFound != -1 && nums2[jFound] < n2 {
				n2NextGreater = n2
				break
			}
		}

		ans[i] = n2NextGreater
	}

	return ans
}
