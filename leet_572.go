package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return checkIsSubtree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkIsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return checkIsSubtree(root.Left, subRoot.Left) && checkIsSubtree(root.Right, subRoot.Right)
}
