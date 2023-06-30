package data_structure

import (
	"fmt"
)

type ll_node struct {
	value int16
	next  *ll_node
}

type linkedlist interface {
	Append(value int16)           // Inserting new node, will be placed at tail
	Prepend(value int16)          // Inserting new node as head node
	Pop()                         // Delete tail node
	PopFirst()                    // Delete head node
	Get(index int16) *ll_node     // Get node of index n from linked list
	Set(index int16, value int16) // Change value of node index n from linked list
	Remove(index int16)           // Remove node of n index from linked list
	Reverse()                     // Reverse the linked list
	Print()
}

type linkedlistImpl struct {
	head   *ll_node
	tail   *ll_node
	length int16
}

func newLinkedList() linkedlist {
	return &linkedlistImpl{}
}

func (ll *linkedlistImpl) Append(value int16) {
	new_node := &ll_node{
		value: value,
	}
	if ll.length == 0 {
		ll.head = new_node
		ll.tail = new_node
	} else {
		ll.tail.next = new_node
		ll.tail = new_node
	}
	ll.length++
}
func (ll *linkedlistImpl) Prepend(value int16) {
	new_node := &ll_node{
		value: value,
	}
	if ll.length == 0 {
		ll.head = new_node
		ll.tail = new_node
	} else {
		temp := ll.head
		ll.head = new_node
		ll.head.next = temp
	}
	ll.length++

}
func (ll *linkedlistImpl) Pop() {
	temp := ll.head
	previous := ll.head
	if ll.length == 1 {
		ll.head = nil
		ll.tail = nil
		ll.length--
		return
	}
	for temp.next != nil {
		previous = temp
		temp = temp.next
	}
	ll.tail = previous
	ll.tail.next = nil
	ll.length--
}

func (ll *linkedlistImpl) PopFirst() {
	temp := ll.head
	if ll.length == 0 {
		ll.head = nil
	} else {
		ll.head = temp.next
		temp = nil
	}
	ll.length--
}

func (ll *linkedlistImpl) Get(index int16) *ll_node {
	if index < 0 || index >= ll.length {
		fmt.Println("Index Out Of Range")
		return nil
	}
	temp := ll.head
	for i := 0; i < int(index); i++ {
		temp = temp.next
	}
	return temp
}

func (ll *linkedlistImpl) Set(index int16, value int16) {
	node := ll.Get(index)
	node.value = value
}

func (ll *linkedlistImpl) Remove(index int16) {
	if index < 0 || index >= ll.length {
		fmt.Println("Index Out Of Range")
		return
	}
	if index == 0 {
		ll.PopFirst()
		return
	}
	if index == int16(ll.length)-1 {
		ll.Pop()
		return
	}
	temp := ll.head
	var previous *ll_node
	for i := 0; i < int(index); i++ {
		previous = temp
		temp = temp.next
	}
	previous.next = temp.next
	temp = nil
	ll.length--
}

func (ll *linkedlistImpl) Reverse() {
	if ll.length == 1 {
		return
	}
	current := ll.head
	var next *ll_node
	var previous *ll_node

	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	ll.head, ll.tail = ll.tail, ll.head
}

func (ll *linkedlistImpl) Print() {
	temp := ll.head
	fmt.Printf("Head |")
	for temp != nil {
		fmt.Printf("%v", temp.value)
		temp = temp.next
		fmt.Printf("|")
	}
	fmt.Printf(" Tail\n")
	fmt.Println("Length: ", ll.length)
	fmt.Println("")
}

/*Singly Linked List*/
func LinkedList() {
	myLinkedList := newLinkedList()
	fmt.Println("Appending 20 and 15")
	myLinkedList.Append(20)
	myLinkedList.Append(15)
	myLinkedList.Print()

	fmt.Println("Prepending 10")
	myLinkedList.Prepend(10)
	myLinkedList.Print()

	fmt.Println("Popping first node, then appending 69 and prepending 420")
	myLinkedList.PopFirst()
	myLinkedList.Append(69)
	myLinkedList.Prepend(420)
	myLinkedList.Print()

	fmt.Println("Get node of index 1 from linked list")
	x := myLinkedList.Get(1)
	fmt.Println(x.value)

	fmt.Println("\nSet value of node index 2 to 66")
	myLinkedList.Set(2, 66)
	myLinkedList.Print()

	fmt.Println("Deleting node index 1 from linked list")
	myLinkedList.Remove(1)
	myLinkedList.Print()

	fmt.Println("Reversing linked list")
	myLinkedList.Reverse()
	myLinkedList.Print()
}
