package queue

import "fmt"

type queue struct {
	items []int
}

// O(1)
func (q *queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// O(1)
func (q *queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// O(1)
func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

// O(1)
func (q *queue) Peek() int {
	return q.items[0]
}

func (q queue) String() string {
	return fmt.Sprintf("%v", q.items)
}

func NewQueue() *queue {
	return &queue{items: make([]int, 0)}
}
