package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//func countNodes(root *TreeNode) int {
//
//	depth := countdeep(root)
//	left := 0;
//	right := 1<<(depth)-1
//
//	rst := binarySearch(root, left, right, depth)
//	return rst
//}
//
//func binarySearch(root *TreeNode,  left int, right int, depth uint) int{
//
//	//if left>right {
//	//	return left-1
//	//}
//	//
//	//mid := (left+right)/2
//	//var cur uint = depth-1
//	//temp := root
//	//for  cur > 0{
//	//	mask := 1<<cur
//	//	cur--
//	//	if mid&mask >0 {
//	//		temp = temp.Right
//	//	} else {
//	//		temp = temp.Left
//	//	}
//	//}
//	//
//	//if temp == nil {
//	//	return binarySearch(root, left, mid-1, depth)
//	//} else {
//	//	return binarySearch(root, mid+1, right, depth)
//	//}
//
//	for left <= right {
//		mid := (left+right)/2
//		cur := depth-1
//		for cur>0
//
//	}
//}
//
//func countdeep(root *TreeNode) uint {
//	var depth uint = 0;
//	for root!=nil {
//		root = root.Left
//		depth++
//	}
//
//	return depth
//}
//func main() {
//	level := 0
//	level++
//	1<<level
//}

