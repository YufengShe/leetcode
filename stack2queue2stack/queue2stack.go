package main

import (
	"fmt"
)

type MyStack struct {
	queue1 []int
	queue2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
	return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue1 = append(this.queue1, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {

	var q, tempq []int

	if len(this.queue1) != 0 {
		q = this.queue1
		tempq = this.queue2
	} else {
		tempq = this.queue1
		q = this.queue2
	}

	for i:=0; i<len(q)-1; i++ {
		tempq = append(tempq, q[i])
	}

	rst := q[len(q)-1]
	q = []int{}

	this.queue1 = tempq
	this.queue2 = q

	return rst
}


/** Get the top element. */
func (this *MyStack) Top() int {
	var q, tempq []int

	if len(this.queue1) != 0 {
		q = this.queue1
		tempq = this.queue2
	} else {
		tempq = this.queue1
		q = this.queue2
	}

	for i:=0; i<len(q)-1; i++ {
		tempq = append(tempq, q[i])
	}

	rst := q[len(q)-1]
	tempq = append(tempq, rst)
	q = []int{}


	this.queue1 = tempq
	this.queue2 = q

	return rst
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.queue1)==0 {
		return true
	} else {
		return false
	}
}