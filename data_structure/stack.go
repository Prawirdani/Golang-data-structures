package data_structure

import "fmt"

/*
	LIFO (Last In First Out)
*/

type stack_node struct {
	value int
	next  *stack_node
}

type stack interface {
	Push(value int) // Insert new node to the stack
	Pop()           // Remove the top node from the Stack
	ShowStack()
}

type stackImpl struct {
	top    *stack_node
	height int
}

func newStack() stack {
	return &stackImpl{}
}

func (s *stackImpl) ShowStack() {
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

func (s *stackImpl) Push(value int) {
	newNode := &stack_node{
		value: value,
	}
	if s.height == 0 {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.height += 1
}

func (s *stackImpl) Pop() {
	if s.top == nil {
		return
	}
	s.top = s.top.next
	s.height -= 1
}

/*Stack, LIFO (Last in First Out)*/
func Stack() {
	myStack := newStack()

	fmt.Println("Pushing 20,30 and 40")
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	myStack.ShowStack()

	fmt.Println("")
	fmt.Println("Popping")
	myStack.Pop()
	myStack.ShowStack()

}
