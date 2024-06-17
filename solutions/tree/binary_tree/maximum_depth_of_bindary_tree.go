package binarytree

func maxDepth(root *TreeNode) int {
	return maxDepthOfBranch(root, 0)
}

func maxDepthOfBranch(node *TreeNode, currentDepth int) int {
	if node == nil {
		return currentDepth
	}

	currentDepth++
	return max(
		maxDepthOfBranch(node.Left, currentDepth),
		maxDepthOfBranch(node.Right, currentDepth),
	)
}
