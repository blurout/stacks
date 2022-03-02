package main

// Stack Node
type StackNode struct {
	data int
	next *StackNode
}

type stack struct {
	top *StackNode
}

func main() {
	// Create new stack
	var MyStack *stack = &stack{top: nil}
	// Add elements
	MyStack.Push(15)
	MyStack.Push(14)
	MyStack.Push(31)
	MyStack.Push(21)
	MyStack.Pop()
	println(MyStack.top.data)
	print(MyStack.Empty())
}

// Add a new element in stack
func (stack *stack) Push(data int) {
	// Make a new stack node
	// And set as top
	stack.top = &StackNode{data: data, next: stack.top}
}

func (stack *stack) Pop() {
	if stack.top == nil {
		print("stack empty")
		return
	}
	stack.top = stack.top.next
}

func (stack *stack) Empty() bool {
	if stack.top == nil {
		return true
	}
	return false
}
