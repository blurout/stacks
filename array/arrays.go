// stack implementation using arrays
package array

import(
	"fmt"
)
type stack struct {
    stack []int
    top int
}

var array stack = stack{top: 0}

func Make_Stack(maxSize int) stack {
	array.stack = make([]int, maxSize)
    return array
}

func (this *stack) Push(x int)  {
    if this.top == len(this.stack) {
        return
    }
	this.stack[this.top] = x
    this.top++
}

func (this *stack) Pop() int {
    if this.top == 0 {
		return -1
	}
	this.top--
    return this.stack[this.top]
}


func (this *stack) Increment(k int, val int)  {
	for i := 0; i < k && i < this.top; i++ {
		this.stack[i] += val
	}
}

func (this stack) Top() int {
	return this.stack[this.top]
}

func (this stack) Print_Stack() {
	fmt.Println("Stack: ")
	for i := this.top; i >= 0; i-- {
		fmt.Println(this.stack[i])
	}
}