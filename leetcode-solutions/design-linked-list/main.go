package main

import "fmt"

type node struct {
	val  int
	prev *node
	next *node
}

type MyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length || this.head == nil {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &node{val: val, next: this.head}

	if this.head != nil {
		this.head.prev = n
	} else {
		this.tail = n
	}

	this.head = n
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &node{val: val, prev: this.tail}
	if this.tail == nil {
		this.head = n
	} else {
		this.tail.next = n
	}
	this.tail = n
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	// If index equals the length of the linked list, the node will be appended to the end of the linked list
	if index == this.length {
		this.AddAtTail(val)
		return
	}

	// If index is greater than the length, the node will not be inserted.
	if index > this.length {
		return
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	n := &node{val: val, next: current, prev: current.prev}
	if current.prev != nil {
		current.prev.next = n
	}
	current.prev = n
	this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length || index < 0 || this.head == nil {
		return
	}

	if index == 0 {
		this.head = this.head.next
		if this.head != nil {
			this.head.prev = nil
		} else {
			this.length = 0
		}
		this.length--
		return
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	if current.next != nil {
		current.next.prev = current.prev
	} else {
		this.tail = current.prev
	}

	if current.prev != nil {
		current.prev.next = current.next
	}

	this.length--
}

func main() {
	l := Constructor()
	l.AddAtHead(13)
	l.AddAtHead(9)
	l.AddAtHead(7)
	l.AddAtIndex(1, 76)
	fmt.Println(l.Get(0))
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
}
