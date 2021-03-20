package main

import "sync"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


var maxSum int
func maxSumBST(root *TreeNode) int {

	maxSum = 0
	isBST(root)
	return maxSum

}


//isBST, min, max
func isBST(node *TreeNode) (ok bool, min int, max int, sum int) {
	if node == nil {
		return true, 4e4, -4e4, 0
	}

	lok, lmin, lmax, lsum := isBST(node.Left)
	rok, rmin, rmax, rsum := isBST(node.Right)

	if lok && rok && (node.Val>lmax) && (node.Val<rmin) {
		sum := lsum+rsum+node.Val
		maxSum = maxf(maxSum,sum)
		return true, minf(lmin, node.Val), maxf(rmax, node.Val), lsum+rsum+node.Val
	} else {
		return false, 0,0,0
	}
}

func maxf(a int, b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}

func minf(a int, b int) int {
	if a<=b{
		return a
	} else {
		return b
	}
}