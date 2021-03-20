package main

import "fmt"

//type Stack struct {
//	S []int
//	Len int
//	Cap int
//}
//
//func factory(length int) *Stack {
//	stack := make([]int, length)
//	return &Stack{stack,
//		0,
//	length,
//	}
//}
//
//func (stack *Stack)push(num int) {
//	if stack.Len == stack.Cap {
//		fmt.Println("Stack is full")
//	}
//	stack.S[stack.Len] = num
//	stack.Len++
//}
//
//
//func (stack *Stack)pop() int{
//	if stack.Len <= 0 {
//		fmt.Println("stack is empty")
//		return 0
//	}
//	num := stack.S[stack.Len-1]
//	stack.Len--
//	return num
//}
//
//func (stack *Stack)getTop() int {
//	if stack.Len <= 0 {
//		fmt.Println("stack is empty")
//		return 0
//	}
//	return stack.S[stack.Len-1]
//}
//
//func (stack *Stack)isEmpty() bool {
//	if stack.Len <= 0{
//		return true
//	} else {
//		return false
//	}
//}
//
//
//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	mp := make(map[int]int)
//	for k,v := range nums1 {
//		mp[v] = k
//	}
//
//	rst := make([]int, len(nums1))
//	stack := factory(len(nums1))
//
//	for i:=len(nums2)-1; i>=0; i-- {
//		for !stack.isEmpty()&&nums2[i]>stack.getTop() {
//			stack.pop()
//		}
//
//		nextGreat := -1
//		if !stack.isEmpty() {
//			nextGreat = stack.getTop()
//		}
//		stack.push(nums2[i])
//
//		if index,ok := mp[nums2[i]]; ok {
//			rst[index] = nextGreat
//		}
//	}
//
//	return rst
//}


func main() {
	a1 := []int{1,3,5,2,4}
	a2 := []int{6,5,4,3,2,1,7}
	fmt.Println(nextGreaterElement(a1, a2))
}