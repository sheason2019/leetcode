package solution110

func pathSum(root *TreeNode, targetSum int) [][]int {
	list := [][]int{}

	if root == nil {
		return list
	}

	recursion(root, targetSum, []int{}, &list)

	return list
}

func recursion(root *TreeNode, targetSum int, buffer []int, list *[][]int) {
	if root == nil {
		return
	}

	nextSum := targetSum - root.Val
	buffer = append(buffer, root.Val)

	if nextSum == 0 && root.Left == nil && root.Right == nil {
		*list = append(*list, append([]int{}, buffer...))
		return
	}

	recursion(root.Left, nextSum, buffer, list)
	recursion(root.Right, nextSum, buffer, list)
}
