package main

import "fmt"

type fifoStack struct {
	stack []string
}

func newFifoStack() *fifoStack {
	return &fifoStack{stack: []string{}}
}

func (s *fifoStack) push(value string) {
	s.stack = append(s.stack, value)
}

func (s *fifoStack) pop() string {
	value := s.stack[0]

	newStack := make([]string, len(s.stack)-1)
	copy(newStack, s.stack[1:])
	s.stack = newStack

	return value
}

func main() {
	myStack := newFifoStack()

	myStack.push("Kirill")
	myStack.push("Mironov")

	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
}
