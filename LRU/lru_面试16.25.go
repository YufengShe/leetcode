package main

import "fmt"

type LRUCache struct {
	mp map[int]*Node
	cap int
	list DoubleList
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	key int
	val int
	pre *Node
	next *Node
}

func (dl *DoubleList)addFirst(node *Node) {


	node.next = dl.head.next
	node.pre = dl.head

	node.next.pre = node
	node.pre.next = node

	dl.size++
}

func (dl *DoubleList)delete(node *Node) {

	node.pre.next = node.next
	node.next.pre = node.pre
	dl.size--
}

func Constructor(capacity int) LRUCache {

	head := new(Node)
	tail := new(Node)

	head.next = tail
	tail.pre = head
	dl := DoubleList{
		head:head,
		tail:tail,
		size:0,
	}

	return LRUCache{
		mp:make(map[int]*Node),
		cap:capacity,
		list:dl,
	}
}





func (this *LRUCache) Get(key int) int {
	if _,ok := this.mp[key];!ok {
		return -1
	}

	node := this.mp[key]
	this.list.delete(node)
	this.list.addFirst(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {

	//key exists
	if _,ok:=this.mp[key];ok {
		node := this.mp[key]
		node.val = value
		this.list.delete(node)
		this.list.addFirst(node)
	} else { // key not existed
		if this.cap == this.list.size {
			node := this.list.tail.pre
			key := node.key
			delete(this.mp, key)
			this.list.delete(node)
		}

		node := &Node{key:key,
			val:value,}

		this.mp[key] = node
		this.list.addFirst(node)
	}
}

type test struct {

}
func main() {

	var a = &test{}
	var b = &test{}
	fmt.Println(a,b)
	fmt.Println(a==b)
}


