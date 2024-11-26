package main

import "fmt"

type lifoStack struct {
	stack []string
}

func newLifoStack() *lifoStack {
	s := make([]string, 0)
	return &lifoStack{stack: s}
}

func (s *lifoStack) push(value string) {
	s.stack = append(s.stack, value)
}

func (s *lifoStack) pop() string {
	tailIndex := len(s.stack) - 1
	value := s.stack[tailIndex]
	s.stack = s.stack[:tailIndex]
	return value
}

func main() {
	myStack := newLifoStack()

	myStack.push("Kirill")
	myStack.push("Mironov")

	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
}
