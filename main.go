package main

import (
	"fmt"

	"github.com/mamadoba/stack-queue/queuestack"
)

func main() {

	var stack = queuestack.Stack{}
	stack.Push(1)
	stack.Push(5)
	stack.Push(9)

	for !stack.IsEmpty() {
		fmt.Println(stack)
		stack.Pop()
	}

	var q = queuestack.Queue{}
	q.Enque(1)
	q.Enque(2)
	q.Enque(4)
	q.Enque(5)

	for !q.IsEmpty() {
		fmt.Println(q)
		q.Deque()
	}
}
