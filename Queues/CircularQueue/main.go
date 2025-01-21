package main

import "fmt"

type Queue struct {
	front int
	rear  int
	size  int
	queue []int
}

func NewCircularQueue(size int) *Queue {
	if size <= 0 {
		return nil
	}

	array := make([]int, size)
	return &Queue{
		front: 0,
		rear:  0,
		queue: array,
		size:  size,
	}
}

func (queue *Queue) isQueueFull() bool {
	if (queue.rear+1)%queue.size == queue.front {
		return true
	}
	return false
}

func (queue *Queue) isQueueEmpty() bool {
	if queue.front == queue.rear {
		return true
	}
	return false
}

func (queue *Queue) Enqueue(x int) {
	if queue.isQueueFull() {
		fmt.Println("Circular Queue is full")
		return
	}

	queue.rear = (queue.rear + 1) % queue.size
	fmt.Println(x, " is inserted at position: ", queue.rear, " front is at: ", queue.front)
	queue.queue[queue.rear] = x
}

func (queue *Queue) Dequeue() int {
	if queue.isQueueEmpty() {
		fmt.Println("Circular queue is empty")
		return 0
	}

	queue.front = (queue.front + 1) % queue.size
	x := queue.queue[queue.front]
	return x
}

func main() {
	queue := NewCircularQueue(7)
	queue.Enqueue(0)
	queue.Enqueue(7)
	queue.Enqueue(8)
	queue.Enqueue(6)
	queue.Enqueue(5)
	queue.Enqueue(1)
	queue.Enqueue(9)

	for i := 0; i < queue.size; i++ {
		fmt.Println("Dequeued element is: ", queue.Dequeue())
	}
}
