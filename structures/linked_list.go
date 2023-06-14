package structures

import (
	"bytes"
	"fmt"
)

type LinkedList struct {
	Top *LinkedListNode
}

type LinkedListNode struct {
	Val  any
	Next *LinkedListNode
}

func NewLinkedList(values ...any) *LinkedList {
	container := &LinkedListNode{}
	current := container
	for _, v := range values {
		current.Next = &LinkedListNode{Val: v}
	}
	return &LinkedList{Top: container.Next}
}

func (l *LinkedList) String() string {
	return l.Top.String()
}

func (l *LinkedListNode) String() string {
	buf := bytes.Buffer{}
	for top := l; top != nil; top = top.Next {
		buf.WriteString(fmt.Sprintf("~%v", top.Val))
	}
	return buf.String()
}
