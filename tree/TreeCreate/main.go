package main

import "fmt"

type Node struct {
	TreeAddress *Tree
	next        *Node
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

func (queue *Queue) Enqueue(address *Tree) {
	n := Node{TreeAddress: address, next: nil}

	if queue.isEmpty() {
		queue.node = &n
		queue.front = queue.node
		queue.rear = queue.node
		return
	}

	queue.rear.next = &n
	queue.rear = queue.rear.next
}

func (queue *Queue) Dequeue() *Tree {
	if queue.isEmpty() {
		fmt.Println("queue is empty")
		return nil
	}

	element := queue.front.TreeAddress
	queue.front = queue.front.next
	queue.node = queue.front
	return element
}

type Tree struct {
	Left  *Tree
	Data  int
	Right *Tree
}

func NewTree() *Tree {
	return &Tree{
		Left:  nil,
		Data:  -1,
		Right: nil,
	}
}

func Create() *Tree {
	fmt.Println("Creating a tree")
	queue := NewQueue()
	root := NewTree()
	var value int
	fmt.Println("Enter the root node")
	fmt.Scanf("%d", &value)
	t := &Tree{Left: nil, Right: nil, Data: value}
	root = t
	queue.Enqueue(root)
	for !queue.isEmpty() {
		p := queue.Dequeue()
		fmt.Println("Enter left child for: ", p.Data)
		fmt.Scanf("%d", &value)
		if value != -1 {
			t := &Tree{Left: nil, Right: nil, Data: value}
			p.Left = t
			queue.Enqueue(t)
		}

		fmt.Println("Enter Right child for: ", p.Data)
		fmt.Scanf("%d", &value)
		if value != -1 {
			t := &Tree{Left: nil, Right: nil, Data: value}
			p.Right = t
			queue.Enqueue(t)
		}
	}

	return root
}

func (t *Tree) RecursivePreOrder() {
	if t != nil {
		fmt.Println(t.Data)
		t.Left.RecursivePreOrder()
		t.Right.RecursivePreOrder()
	}
}

func (t *Tree) RecursiveInOrder() {
	if t != nil {
		t.Left.RecursivePreOrder()
		fmt.Println(t.Data)
		t.Right.RecursivePreOrder()
	}
}

func (t *Tree) RecursivePostOrder() {
	if t != nil {
		t.Left.RecursivePreOrder()
		t.Right.RecursivePreOrder()
		fmt.Println(t.Data)
	}
}

func (t *Tree) LevelOrderTraversal() {
	q := NewQueue()
	q.Enqueue(t)
	for !q.isEmpty() {
		p := q.Dequeue()
		fmt.Println(p.Data)
		if p.Left != nil {
			q.Enqueue(p.Left)
		}
		if p.Right != nil {
			q.Enqueue(p.Right)
		}
	}
}

func (t *Tree) CountNodesWithRecursion() int {
	if t != nil {
		x := t.Left.CountNodesWithRecursion()
		y := t.Right.CountNodesWithRecursion()
		return x + y + 1

	}
	return 0
}

func (t *Tree) CountNodesWithoutRecursion() int {
	q := NewQueue()
	q.Enqueue(t)
	count := 0
	for !q.isEmpty() {
		p := q.Dequeue()
		count++
		if p.Left != nil {
			q.Enqueue(p.Left)
		}

		if p.Right != nil {
			q.Enqueue(p.Right)
		}
	}
	return count
}

func (t *Tree) Height() int {
	if t != nil {
		x := t.Left.Height()
		y := t.Right.Height()
		if x > y {
			return x + 1
		} else {

		}
		return y + 1

	}
	return 0
}

func main() {
	t := Create()
	fmt.Println("Pre Order")
	t.RecursivePreOrder()
	fmt.Println("In Order")
	t.RecursiveInOrder()
	fmt.Println("Post Order")
	t.RecursivePostOrder()
	fmt.Println("Level Order")
	t.LevelOrderTraversal()
	fmt.Println("Count nodes with recursion")
	fmt.Println(t.CountNodesWithRecursion())
	fmt.Println("Count nodes without recursion")
	fmt.Println(t.CountNodesWithoutRecursion())
	fmt.Println("Height")
	fmt.Println(t.Height())
}
