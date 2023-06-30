package data_structure

import "fmt"

type queue_node struct {
	value int
	next  *queue_node
}

type queue interface {
	Enqueue(value int) // Inserting new node to queue
	Dequeue()          // Dequeuing
	ShowQueue()
}

type queueImpl struct {
	front  *queue_node
	rear   *queue_node
	length int
}

func newQueue() queue {
	return &queueImpl{}
}

// Enqueue method for enqueuing a new element
func (q *queueImpl) Enqueue(value int) {
	newNode := &queue_node{
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
func (q *queueImpl) Dequeue() {
	if q.front == nil {
		return
	}
	temp := q.front
	q.front = temp.next
	temp = nil
	q.length -= 1
}

func (q *queueImpl) ShowQueue() {
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
	fmt.Println("")
}

/*Queue, FIFO (First in First Out)*/
func Queue() {
	myQueue := newQueue()
	fmt.Println("Enqueuing 10, 15 and 20")
	myQueue.Enqueue(10)
	myQueue.Enqueue(15)
	myQueue.Enqueue(20)
	myQueue.ShowQueue()

	fmt.Println("Dequeuing")
	myQueue.Dequeue()
	myQueue.ShowQueue()
}
