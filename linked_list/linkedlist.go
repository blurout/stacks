package linked_list

import(
	"fmt"
)
// Stack Node
type StackNode struct {
	data int
	next *StackNode
}

type linked_list_stack struct {
	top *StackNode
}

func Make_Stack() linked_list_stack {
    var mystack linked_list_stack = linked_list_stack{top: nil}
    return mystack
}
var MyStack *linked_list_stack = &linked_list_stack{top: nil}

// Add a new element in stack
func (stack *linked_list_stack) Push(data int) {
	// Make a new stack node
	// And set as top
	stack.top = &StackNode{data: data, next: stack.top}
}

func (stack *linked_list_stack) Pop() {
	if stack.top == nil {
		print("stack empty")
		return
	}
	stack.top = stack.top.next
}

func (stack *linked_list_stack) Empty() bool {
	if stack.top == nil {
		return true
	}
	return false
}

func (stack *linked_list_stack) Peek() int {
	if !stack.Empty() {
		return stack.top.data
	}
	fmt.Println("error: stack empty")
	return -1
}

func (stack *linked_list_stack) Print() {
	tmp := stack.top
	fmt.Println("current stack: ")
	for tmp != nil {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}
