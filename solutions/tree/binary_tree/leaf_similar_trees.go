package binarytree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i, val := range leaves1 {
		if leaves2[i] != val {
			return false
		}
	}

	return true
}

func getLeaves(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	return append(getLeaves(node.Left), getLeaves(node.Right)...)
}
