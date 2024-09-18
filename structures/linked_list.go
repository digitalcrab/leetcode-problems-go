package structures

import (
	"bytes"
	"fmt"
)

func NewSinglyLikedList[T comparable](elements ...T) *SinglyLikedList[T] {
	if len(elements) == 0 {
		return &SinglyLikedList[T]{}
	}

	list := &SinglyLikedList[T]{
		Head: &SinglyLikedListNode[T]{
			Value: elements[0],
		},
	}

	// point the current one to the head
	current := list.Head

	// start looping over the list from the second element
	for _, e := range elements[1:] {
		// new node
		node := &SinglyLikedListNode[T]{Value: e}
		// while current still the same assign it's next to a new node and
		// switch current to the new node as well
		current.Next, current = node, node
	}

	return list
}

// SinglyLikedListNode holds the value and a link to the next element
// as we use generics we have to make sure at least types are comparable
type SinglyLikedListNode[T comparable] struct {
	Value T
	Next  *SinglyLikedListNode[T]
}

// SinglyLikedList holds the lint to the first node of the list
type SinglyLikedList[T comparable] struct {
	Head *SinglyLikedListNode[T]
}

func (ll *SinglyLikedList[T]) Display() string {
	buf := new(bytes.Buffer)

	// iterate over the linked list and write the values to the buffer
	// current is going to represent the head initially, and we iterate until
	// current is not equal to nil, moving current to the next element
	for current := ll.Head; current != nil; current = current.Next {
		_, err := fmt.Fprintf(buf, "%v -> ", current.Value)
		if err != nil {
			return "<error>"
		}
	}

	// at the list is going to have nil
	buf.WriteString("nil")

	return buf.String()
}

// InsertAtBeginning ass an element to the beginning of a list
// Time complexity is O(1) as we have just 2 operations (create node, reassign head)
func (ll *SinglyLikedList[T]) InsertAtBeginning(v T) {
	// create a new list node that holds a new value and set the next
	// pointer to the current head of the list
	node := &SinglyLikedListNode[T]{
		Value: v,
		Next:  ll.Head,
	}

	// assign head to the newly created node
	ll.Head = node
}

func (ll *SinglyLikedList[T]) InsertAtEnd(v T) {
	node := &SinglyLikedListNode[T]{
		Value: v,
	}

	// case when list is still empty, we simply assign head to the new node
	if ll.Head == nil {
		ll.Head = node
		return
	}

	// move to the very end of the list
	current := ll.Head

	// ! important that we move until the current.next is not nil !
	for ; current.Next != nil; current = current.Next {
	}

	// when we found the very last element we set it's next pointer to the newly created node
	current.Next = node
}

// Search tells if v is in the list, time complexity is O(n) as we iterate over all elements
func (ll *SinglyLikedList[T]) Search(v T) bool {
	for current := ll.Head; current != nil; current = current.Next {
		if current.Value == v {
			return true
		}
	}

	return false
}
