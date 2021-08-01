package main

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// after the first once over if none of the above are right then
	// do a recursive call on the left side first then the right side
	leftSide := minDepth(root.Left)
	rightSide := minDepth(root.Right)

	return 1 + height(leftSide, rightSide)

}

func height(a, b int) int {
	if a <= b {
		return a + 1
	}
	return b + 1
}
