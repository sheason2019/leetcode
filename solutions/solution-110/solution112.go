package solution110

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nextSum := targetSum - root.Val

	if nextSum == 0 && (root.Left == nil && root.Right == nil) {
		return true
	}

	return hasPathSum(root.Left, nextSum) || hasPathSum(root.Right, nextSum)
}
