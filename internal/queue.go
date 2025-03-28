package internal

import "fmt"

// Queue is a simple FIFO queue implementation
// item[0] (=HEAD) - item[1] ... - item[n-1] (=TAIL)
type Queue struct {
	items    []*Node
	size     int
	capacity int
}

// NewQueue creates a new Queue with the specified maximum capacity
func NewQueue(maxCapacity int) *Queue {
	return &Queue{
		items:    make([]*Node, 0, maxCapacity),
		size:     0,
		capacity: maxCapacity,
	}
}

// Enqueue adds an item to the queue if it's not full
func (q *Queue) Enqueue(item *Node) error {
	if len(q.items) >= q.capacity {
		return fmt.Errorf("queue is full")
	}
	q.items = append(q.items, item)
	q.size++
	return nil
}

// Dequeue removes and returns the front item from the queue
func (q *Queue) Dequeue() *Node {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

// Full checks if the queue is full
func (q *Queue) Full() bool {
	return len(q.items) >= q.capacity
}

// Empty checks if the queue is empty
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Tail() *Node {
	if q.Empty() {
		return nil
	}
	return q.items[len(q.items)-1]
}

func (q *Queue) Head() *Node {
	if q.Empty() {
		return nil
	}
	return q.items[0]
}
