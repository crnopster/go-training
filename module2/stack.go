package main

// Stack .
type Stack struct {
	top   *node
	count int
}
type node struct {
	value float64
	next  *node
}

// NewStack .
func NewStack() *Stack {
	stack := new(Stack)
	return stack
}

// Push elements in the Stack
func (stack *Stack) Push(value float64) {
	n := new(node)
	n.value = value
	n.next = stack.top

	stack.top = n
	stack.count = stack.count + 1
}

// Size return nubmer of items in Stack
func (stack *Stack) Size() int {
	return stack.count
}

// Pop elements from the Stack
func (stack *Stack) Pop() float64 {
	v := stack.top.value
	stack.top = stack.top.next
	stack.count = stack.count - 1
	return v
}

// Peek return the last element added to the Stack
func (stack *Stack) Peek() float64 {
	return stack.top.value
}

// Contains check if value exist in the Stack
func (stack Stack) Contains(value float64) bool {
	for i := 0; i < stack.count; i++ {
		if stack.top.value == value {
			return true
		} else if stack.top.value != value {
			stack.top = stack.top.next
		}
	}
	return false
}

// IsEmpty check if Stack is empty
func (stack *Stack) IsEmpty() bool {
	if stack.count > 0 {
		return false
	}
	return true
}
