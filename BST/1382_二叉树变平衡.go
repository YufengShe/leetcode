package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {

	arr := make([]int,0)
	var track func(node *TreeNode)

	track = func(node *TreeNode) {
		if node == nil {
			return
		}

		track(node.Left)
		arr = append(arr, node.Val)
		track(node.Right)
	}

	track(root)

	root = buildBTree(arr, 0, len(arr)-1)
	return root

}

func buildBTree(arr []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left+right)/2
	node := &TreeNode{
		Val:arr[mid],
	}

	node.Left = buildBTree(arr, left, mid-1)
	node.Right = buildBTree(arr, mid+1, right)

	return node
}


