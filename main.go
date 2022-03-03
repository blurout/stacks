package main

import (
	"stacks/linked_list"
	"stacks/array"
)

func main() {

	linked_list.MyStack.Push(15)
	linked_list.MyStack.Print()

	array.Mystack.Push(23)
	array.Mystack.Push(3)
	array.Mystack.Pop()
	array.Mystack.Print_Stack()
}
