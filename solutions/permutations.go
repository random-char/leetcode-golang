package solutions

func Permute(nums []int) [][]int {
	return permute(nums)
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	heapPermute(&result, nums, len(nums))
	return result
}

func heapPermute(result *[][]int, nums []int, size int) {
	if size == 1 {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}

	for i := 0; i < size; i++ {
		heapPermute(result, nums, size-1)

		if size%2 == 1 {
			swap(nums, 0, size-1)
		} else {
			swap(nums, i, size-1)
		}
	}
}

func swap(nums []int, i, j int) {
	tmp := (nums)[i]
	(nums)[i] = (nums)[j]
	(nums)[j] = tmp
}
