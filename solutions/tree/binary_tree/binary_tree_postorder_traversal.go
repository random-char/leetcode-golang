package binarytree

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
	return append(result, root.Val)
}
