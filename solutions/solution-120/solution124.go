package solution120

func maxPathSum(root *TreeNode) int {
	maxValue := root.Val
	maxPathSumRecursion(root, &maxValue)

	return maxValue
}

func maxPathSumRecursion(root *TreeNode, maxValue *int) int {
	if root == nil {
		return 0
	}

	leftMPS := maxPathSumRecursion(root.Left, maxValue)
	if leftMPS < 0 {
		leftMPS = 0
	}
	rightMPS := maxPathSumRecursion(root.Right, maxValue)
	if rightMPS < 0 {
		rightMPS = 0
	}

	// 该节点能获取到的最大值
	currentMaxVal := leftMPS + rightMPS + root.Val
	if currentMaxVal > *maxValue {
		*maxValue = currentMaxVal
	}

	// 该节点能产生的最大贡献值
	currentContributionValue := root.Val
	temp := currentContributionValue + leftMPS
	tempRight := currentContributionValue + rightMPS
	if tempRight > temp {
		temp = tempRight
	}
	if temp > currentContributionValue {
		currentContributionValue = temp
	}

	return currentContributionValue
}
