package main
// Go program for
// Implementation stack using linked list

// Stack Node
type StackNode struct {
	data int
	next * StackNode
}
func getStackNode(data int, top * StackNode) * StackNode {
	// return new StackNode
	return &StackNode {
		data,
		top,
	}
}
type MyStack struct {
	top * StackNode
	count int
}
// Add a new element in stack
func(this *MyStack) push(data int) {
	// Make a new stack node
	// And set as top
	this.top = getStackNode(data, this.top)
	// Increase node value
	this.count++
}

func main() {
	// Create new stack 
	var s * MyStack = &MyStack{top: nil, count: 0}
	// Add element
	s.push(15)
	s.push(14)
	s.push(31)
	s.push(21)
	print(s.top.data)
}