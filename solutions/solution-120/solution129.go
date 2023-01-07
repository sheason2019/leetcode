package solution120

func sumNumbers(root *TreeNode) int {
	sum := 0

	sumNumbersRecursion(root, 0, &sum)

	return sum
}

func sumNumbersRecursion(node *TreeNode, base int, sum *int) {
	val := base*10 + node.Val

	if node.Left == nil && node.Right == nil {
		*sum = *sum + val
		return
	}

	if node.Left != nil {
		sumNumbersRecursion(node.Left, val, sum)
	}
	if node.Right != nil {
		sumNumbersRecursion(node.Right, val, sum)
	}
}
