package main

import (
	"stacks/linked_list"
	"stacks/array"
)

func main() {

	stack := linked_list.Make_Stack()
	stack.Push(6)
	stack.Pop()
	stack.Push(5)
	if stack.Empty() {
		println("stack empty")
	} else {
		stack.Print()
	}

	array_stack := array.Make_Stack(10)
	array_stack.Push(23)
	array_stack.Push(3)
	array_stack.Pop()
	array_stack.Print_Stack()
}
