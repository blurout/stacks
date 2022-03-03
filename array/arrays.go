// stack implementation using arrays
package array

import(
	"fmt"
)
type stack [MAX_SIZE]int
const MAX_SIZE int = 101
var top int = -1

var Mystack stack

func (array stack) Push(n int) {
	if top == MAX_SIZE - 1 {
		fmt.Println("Error: Stack Overflow")
		return
	}
	top++
	Mystack[top] = n
}

func (array stack) Pop() {
	if top == - 1 {
		fmt.Println("Error: Stack empty, No element to pop")
		return
	}
	top--
}

func (array stack) Top() int {
	return array[top]
}

func (array stack) Print_Stack() {
	fmt.Println("Stack: ")
	for i := top; i >= 0; i-- {
		fmt.Println(Mystack[i])
	}
}
