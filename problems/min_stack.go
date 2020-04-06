package problems

import "math"

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// push(x) -- Push element x onto stack.
// pop() -- Removes the element on top of the stack.
// top() -- Get the top element.
// getMin() -- Retrieve the minimum element in the stack.
//
// Example:
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> Returns -3.
// minStack.pop();
// minStack.top();      --> Returns 0.
// minStack.getMin();   --> Returns -2.

type minStackNode struct {
	prev *minStackNode
	val  int
	min  float64
}

type MinStack struct {
	top *minStackNode
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (m *MinStack) Push(x int) {
	// we can safely assume that stack won't change below the new value
	// so we calculate min value on every push and store it on the same level
	min := float64(x)
	if m.top != nil {
		min = math.Min(m.top.min, min)
	}

	m.top = &minStackNode{
		prev: m.top,
		val:  x,
		min:  min,
	}
}

func (m *MinStack) Pop() {
	if m.top != nil {
		m.top = m.top.prev
	}
}

func (m *MinStack) Top() int {
	if m.top != nil {
		return m.top.val
	}
	return 0
}

func (m *MinStack) GetMin() int {
	if m.top != nil {
		return int(m.top.min)
	}
	return 0
}
