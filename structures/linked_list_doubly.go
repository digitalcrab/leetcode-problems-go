package structures

import (
	"bytes"
	"fmt"
	"iter"
)

func NewDoublyLikedList[T comparable](elements ...T) *DoublyLikedList[T] {
	if len(elements) == 0 {
		return &DoublyLikedList[T]{}
	}

	first := &DoublyLikedListNode[T]{Value: elements[0]}

	// create a list with a first element as head and tail
	list := &DoublyLikedList[T]{
		Head: first,
		Tail: first,
	}

	current := list.Head

	// loop over the rest
	for _, e := range elements[1:] {
		// create a new node where prev assigned to the current
		node := &DoublyLikedListNode[T]{Value: e, Prev: current}
		// replace current and next of the current with a new node
		current.Next, current = node, node
		// update the tail
		list.Tail = node
	}

	return list
}

type DoublyLikedListNode[T comparable] struct {
	Value T
	Prev  *DoublyLikedListNode[T]
	Next  *DoublyLikedListNode[T]
}

type DoublyLikedList[T comparable] struct {
	Head *DoublyLikedListNode[T]
	Tail *DoublyLikedListNode[T]
}

func (ll *DoublyLikedList[T]) Display(reverse bool) string {
	buf := new(bytes.Buffer)

	for v := range ll.items(reverse) {
		_, err := fmt.Fprintf(buf, "%v -> ", v)
		if err != nil {
			return "<error>"
		}
	}

	// at the list is going to have nil
	buf.WriteString("nil")

	return buf.String()
}

func (ll *DoublyLikedList[T]) items(reverse bool) iter.Seq[T] {
	if reverse {
		return ll.itemsReverse
	}
	return ll.itemsDirect
}

// itemsDirect is an iterator implementation in direct order
func (ll *DoublyLikedList[T]) itemsDirect(yield func(T) bool) {
	// current is going to represent the head initially, and we iterate until
	// current is not equal to nil, moving current to the next element
	for current := ll.Head; current != nil; current = current.Next {
		if !yield(current.Value) {
			return
		}
	}
}

// itemsReverse is an iterator implementation in reverse order
func (ll *DoublyLikedList[T]) itemsReverse(yield func(T) bool) {
	// current is going to represent the tail initially, and we iterate until
	// current is not equal to nil, moving current to the prev element
	for current := ll.Tail; current != nil; current = current.Prev {
		if !yield(current.Value) {
			return
		}
	}
}

func (ll *DoublyLikedList[T]) InsertAtBeginning(v T) {
	// create a new list node that holds a new value
	node := &DoublyLikedListNode[T]{Value: v}

	// case when list is still empty, we simply assign head to the new node
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	// set the next pointer to the current head of the list
	node.Next = ll.Head

	// assign prev element of the current head to the newly created node
	ll.Head.Prev = node

	// replace the head with a new node
	ll.Head = node
}

func (ll *DoublyLikedList[T]) InsertAtEnd(v T) {
	node := &DoublyLikedListNode[T]{Value: v}

	// case when list is still empty, we simply assign head to the new node
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	// prev element of a new node should be equal to the tail of a list
	node.Prev = ll.Tail

	// next element of a tail should be a newly created node
	ll.Tail.Next = node

	// replace the tail with a new node
	ll.Tail = node
}
