package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	//add new head
	nHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	nHead.Next = head
	tempHead := nHead
	left++ //>=2
	right++ //>=2
	count := 1


	//get prev, the previous node of left
	var prev *ListNode
	for tempHead!=nil && count<left{
		if count == left-1 {
			prev = tempHead
		}

		tempHead = tempHead.Next
		count++
	}


	leftNode := tempHead

	for count<=right{
		temp := tempHead.Next
		tempHead.Next = prev.Next
		prev.Next = tempHead

		tempHead = temp
		count++
	}

	leftNode.Next = tempHead


	return nHead.Next


}