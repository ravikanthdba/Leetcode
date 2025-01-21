package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	node  *Node
	front *Node
	rear  *Node
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
		node:  nil,
	}
}

func (queue *Queue) isEmpty() bool {
	return queue.node == nil
}

func (queue *Queue) Enqueue(value int) {
	n := Node{value: value, next: nil}

	if queue.isEmpty() {
		queue.node = &n
		queue.front = queue.node
		queue.rear = queue.node
		return
	}

	queue.rear.next = &n
	queue.rear = queue.rear.next
}

func (queue *Queue) Display() {
	current := queue.front

	for current != queue.rear {
		fmt.Println("Value is: ", current.value)
		current = current.next
	}

	fmt.Println("Value is: ", current.value)

}

func (queue *Queue) Dequeue() {
	if queue.isEmpty() {
		fmt.Println("queue is empty")
		return
	}

	element := queue.front.value
	queue.front = queue.front.next
	queue.node = queue.front
	fmt.Println("returned element is: ", element)
}

func main() {
	q := NewQueue()
	q.Enqueue(3)
	q.Enqueue(6)
	q.Enqueue(9)
	q.Enqueue(12)
	q.Enqueue(15)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q.isEmpty())

	q.Enqueue(4)
	q.Enqueue(8)
	q.Enqueue(12)
	q.Enqueue(16)
	q.Enqueue(20)
	fmt.Println(q.isEmpty())

	q.Display()

}
