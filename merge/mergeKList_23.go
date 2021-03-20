package main
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoList(la *ListNode, lb *ListNode) *ListNode {

	if la == nil {
		return lb
	} else if lb == nil{
		return la
	}

	head := new(ListNode)
	tail := head
	for la!=nil && lb!=nil {
		if la.Val <= lb.Val {
			tail.Next = la
			la = la.Next
		} else {
			tail.Next = lb
			lb = lb.Next
		}
		tail =tail.Next
	}

	if la != nil {
		tail.Next = la
	} else if lb != nil {
		tail.Next = lb
	}

	return head.Next

}


func merge(lists []*ListNode, left int, right int) *ListNode {
	if left==right {
		return lists[left]
	}

	if left>right { //针对的是len(lists)-1=-1的情况
		return nil
	}

	mid := (left+right)/2
	return mergeTwoList(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}



