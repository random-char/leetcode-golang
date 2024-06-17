package binarytree

func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2
	root := TreeNode{Val: nums[mid]}

	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid-1])
	}

	if mid < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return &root
}
