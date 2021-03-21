package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var ans *TreeNode

	var backOrder func(node *TreeNode) bool

	backOrder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		isLeft := backOrder(node.Left)
		isRight := backOrder(node.Right)

		if (isLeft && isRight) || ((isLeft || isRight)&&(node==p||node==q)) {
			ans = node
			return true
		}

		return isLeft || isRight || node==p || node==q
	}

	backOrder(root)
	return ans
}