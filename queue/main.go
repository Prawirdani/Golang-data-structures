package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front  *Node
	rear   *Node
	length int
}

func InitQueue(value int) *Queue {
	newQueue := &Queue{
		front: &Node{
			value: value,
		},
		rear: &Node{
			value: value,
		},
		length: 1,
	}
	return newQueue
}

// Enqueue method for enqueuing a new element
func (q *Queue) Enqueue(value int) {
	newNode := &Node{
		value: value,
	}
	if q.length == 0 {
		q.front = newNode
		q.rear = newNode
	} else if q.length == 1 {
		q.front.next = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.length++
}

// Dequeue method for dequeuing the first element or the front element of the queue
func (q *Queue) Dequeue() {
	if q.front == nil {
		return
	}
	temp := q.front
	q.front = temp.next
	temp = nil
	q.length -= 1
}

func (q *Queue) PrintQueue() {
	if q.length == 0 {
		fmt.Println("Empty queue")
		return
	}
	temp := q.front
	for temp != nil {
		fmt.Printf("| %d ", temp.value)
		temp = temp.next
	}
	fmt.Println("|")
	fmt.Printf("Length\t: %d\n", q.length)
	fmt.Printf("Front\t: %d\n", q.front.value)
	fmt.Printf("Rear\t: %d\n", q.rear.value)
}

func main() {
	myQueue := InitQueue(5)
	myQueue.Enqueue(10)
	myQueue.Enqueue(15)
	myQueue.Enqueue(20)
	myQueue.PrintQueue()

	fmt.Println("")
	myQueue.Dequeue()
	myQueue.Dequeue()
	myQueue.PrintQueue()
}
