package binarytree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 1)
	result[0] = root.Val

	result = append(result, preorderTraversal(root.Left)...)
	return append(result, preorderTraversal(root.Right)...)
}
