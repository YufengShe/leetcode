package heap
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeKLists(lists []*ListNode) *ListNode {

	var head *ListNode
	var tail *ListNode

	//build small top heap
	heap := make([]*ListNode, len(lists))

	j:=0
	for i:=0; i<len(lists); i++{
		if lists[i] == nil {
			continue
		} else {
			heap[j] = lists[i]
			j++
		}

	}
	length := j


	if length <= 0 {
		return nil
	}

	for i:=length/2-1; i>=0; i-- {
		heapAdjust(heap, i, length)
	}

	head = heap[0]
	tail = heap[0]

	for length>0 {
		heap[0] = heap[0].Next
		if heap[0] == nil {
			heap[0] = heap[length-1]
			length--
		}
		heapAdjust(heap, 0, length)
		tail.Next = heap[0]
		tail = heap[0]
	}

	return head
}

func heapAdjust(heap []*ListNode, i int, length int) {

	min := i
	left := 2*i+1
	right := 2*i+2

	if left<length && heap[left].Val<heap[min].Val {
		min = left
	}

	if right<length && heap[right].Val<heap[min].Val {
		min = right
	}

	if min != i {
		heap[min], heap[i] = heap[i], heap[min]
		heapAdjust(heap, min, length)
	}

}
