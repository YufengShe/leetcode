package main


//recursive
//func reverseList(head *ListNode) *ListNode {
//
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	last := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//
//	return last
//
//}


//for loop

func reverseList(head *ListNode) *ListNode {


	if head == nil || head.Next == nil {
		return head
	}

	nHead := &ListNode{
		Val:0,
		Next: head,
	}


	first := head

	for head!=nil {
		temp := head.Next

		head.Next = nHead.Next
		nHead.Next = head

		head =temp
	}

	first.Next = nil

	return nHead.Next

}