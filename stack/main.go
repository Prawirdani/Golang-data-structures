package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	height int
}

func InitStack(value int) *Stack {
	newStack := &Stack{
		top: &Node{
			value: value,
		},
		height: 1,
	}
	return newStack
}

func (s *Stack) PrintStack() {
	if s.height == 0 {
		fmt.Println("Empty Stack")
		return
	}
	temp := s.top
	for temp != nil {
		if temp.value == s.top.value {
			fmt.Printf("| %d | <-- Top\n", temp.value)
		} else {
			fmt.Printf("| %d |\n", temp.value)
		}
		temp = temp.next
	}
	fmt.Printf("Stack Height: %d\n", s.height)
}

// Push Method for Add a Node to the stack
func (s *Stack) Push(value int) {
	newNode := &Node{
		value: value,
	}
	if s.height == 0 {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
		s.height += 1
	}
}

// Pop Method for removing the Top Node from the Stack
func (s *Stack) Pop() {
	if s.top == nil {
		return
	}
	s.top = s.top.next
	s.height -= 1
}

func main() {
	myStack := InitStack(10)
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	myStack.PrintStack()

	fmt.Println("")
	myStack.Pop()

	myStack.PrintStack()
}
