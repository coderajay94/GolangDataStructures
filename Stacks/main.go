package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if size := len(s.items); size > 0 {
		lastItem := s.items[size-1]
		s.items = append(s.items[:size-1])
		return lastItem
	}
	return -1
}

func (s *Stack) Peek() int {
	if size := len(s.items); size > 0 {
		return s.items[size-1]
	}
	return -1
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {

	fmt.Println("testing with stack in golang")

	stack := Stack{items: []int{12, 123, 133}}

	stack.Push(1233)
	stack.Push(33)

	fmt.Println("printing the stack before popping", stack)
	lastElement := stack.Pop()
	fmt.Println("popped last element", lastElement)
	fmt.Println("printing the stack after popping", stack)

	fmt.Print("last element present now", stack.Peek())

	fmt.Println("length of stack is ", stack.Size())
	fmt.Println("length of stack is empty", stack.IsEmpty())
	stack.Pop()
	fmt.Println("length of stack is ", stack.Size())
	stack.Pop()
	fmt.Println("length of stack is ", stack.Size())
	stack.Pop()
	fmt.Println("length of stack is ", stack.Size())
	stack.Pop()
	fmt.Println("length of stack is ", stack.Size())
	stack.Pop()
	fmt.Println("length of stack is ", stack.Size())
	fmt.Println("length of stack is empty", stack.IsEmpty())
}
