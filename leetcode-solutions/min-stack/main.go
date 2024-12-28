package main

import (
	"fmt"
)

type node struct {
	val int
	min int // every node is going to have current min
}

type MinStack struct {
	stack []*node // in the stack we store not just value but additional info
}

func Constructor() MinStack {
	return MinStack{}
}

// Push pushes the element val onto the stack.
func (s *MinStack) Push(val int) {
	currentMin := val

	// if there is something inside, then we use this as current min
	if s.top() != nil {
		currentMin = s.top().min
	}

	n := &node{
		val: val,
		min: min(val, currentMin), // with a new node we store current min
	}

	s.stack = append(s.stack, n)
}

// Pop removes the element on the top of the stack.
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1] // [:), means last not included
}

// Top gets the top element of the stack.
func (s *MinStack) Top() int {
	if s.top() == nil {
		panic("should not happen based on constrains")
	}
	return s.top().val
}

func (s *MinStack) top() *node {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

// GetMin retrieves the minimum element in the stack.
func (s *MinStack) GetMin() int {
	return s.top().min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
