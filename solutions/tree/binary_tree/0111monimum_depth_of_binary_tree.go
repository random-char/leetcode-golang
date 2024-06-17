package binarytree

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return minDepthRec(root)
}

func minDepthRec(node *TreeNode) int {
	if node == nil {
		return math.MaxInt
	}

	if node.Right == nil && node.Left == nil {
		return 1
	}

	return min(minDepth(node.Left), minDepth(node.Right)) + 1
}
