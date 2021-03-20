package main

import (
	"container/heap"
	"fmt"
)

//type intHeap []int
//
//func (h intHeap)Len() int{
//	return len(h)
//}
//
//func (h intHeap)Swap(i int, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h intHeap)Less(i int, j int) bool {
//	return h[i]>h[j]  //小顶堆
//}
//
//func (h *intHeap)Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//
//func (h *intHeap)Pop() interface{} {
//
//	rst := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return rst
//}

type elem struct {
	val int
	index int
}
type eHeap []*elem

//Len() Less() Swap() Push() Pop()

func (h eHeap)Len() int{
	return len(h)
}

func (h eHeap)Swap(i int, j int)  {
	h[i], h[j] = h[j], h[i]
}

func (h *eHeap)Less(i int, j int) bool {
	return (*h)[i].val > (*h)[j].val
}

func (h *eHeap)Push(x interface{})  {
	*h = append(*h, x.(*elem))
}

func (h *eHeap)Pop() interface{} {

	rst := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return rst
}

func maxSlidingWindow(nums []int, k int) []int {

	rst := make([]int, 0)
	h := new(eHeap)
	for i:=0; i<k; i++ {
		e := &elem{
			val:nums[i],
			index:i,
		}
		heap.Push(h, e)
	}

	rst = append(rst, (*h)[0].val)
	for i:=k; i<len(nums); i++ {
		e := &elem{
			val:nums[i],
			index:i,
		}
		heap.Push(h, e)

		for (*h)[0].index <= i-k {
			heap.Pop(h)
		}

		rst = append(rst, (*h)[0].val)
	}

	return rst
}

