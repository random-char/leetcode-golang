package binarytree

func isSymmetric(root *TreeNode) bool {
	return areSubtreesSymmetric(root.Right, root.Left)
}

func areSubtreesSymmetric(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}

	return l.Val == r.Val && areSubtreesSymmetric(l.Right, r.Left) && areSubtreesSymmetric(l.Left, r.Right)
}
