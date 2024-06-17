package binarytree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	return recurBinTreePaths(root, "")
}

func recurBinTreePaths(node *TreeNode, currentPath string) []string {
	if len(currentPath) != 0 {
		currentPath += "->"
	}
	currentPath += fmt.Sprintf("%d", node.Val)

	result := make([]string, 0)
	if node.Left == nil && node.Right == nil {
		result = append(result, currentPath)
	} else {
		if node.Left != nil {
			result = append(result, recurBinTreePaths(node.Left, currentPath)...)
		}

		if node.Right != nil {
			result = append(result, recurBinTreePaths(node.Right, currentPath)...)
		}

	}

	return result
}
