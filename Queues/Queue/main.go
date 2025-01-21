package main

import "fmt"

type Queue struct {
	front int
	rear  int
	size  int
	queue []int
}

func NewQueue(size int) *Queue {
	if size <= 0 {
		fmt.Println("Queue cannot be created")
		return nil
	}
	return &Queue{
		size:  size,
		front: -1,
		rear:  -1,
		queue: make([]int, size), // Preallocate a fixed-size slice
	}
}

func (q *Queue) isEmpty() bool {
	return q.front == -1
}

func (q *Queue) isFull() bool {
	return q.rear == q.size-1
}

func (q *Queue) Enqueue(value int) {
	if q.isFull() {
		fmt.Println("Queue is full, cannot insert")
		return
	}

	if q.front == -1 {
		q.front = 0 // Set front to 0 for the first element
	}
	q.rear++
	q.queue[q.rear] = value
}

func (q *Queue) Dequeue() {
	if q.isEmpty() {
		fmt.Println("Queue is empty, cannot delete")
		return
	}

	// Retrieve the value at the front
	element := q.queue[q.front]
	fmt.Println("Returned element is:", element)

	// Move front pointer forward
	if q.front == q.rear {
		// Queue is empty after this dequeue
		q.front, q.rear = -1, -1
	} else {
		q.front++
	}
}

func main() {
	queue := NewQueue(7)
	queue.Enqueue(0)
	queue.Enqueue(7)
	queue.Enqueue(8)
	queue.Enqueue(6)
	queue.Enqueue(5)
	queue.Enqueue(1)
	queue.Enqueue(9)

	// Attempt to enqueue to a full queue
	queue.Enqueue(10)

	// Dequeue all elements
	for i := 0; i < 7; i++ {
		queue.Dequeue()
	}

	// Attempt to dequeue from an empty queue
	queue.Dequeue()
}
