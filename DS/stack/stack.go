package stack

import "fmt"

type stack struct {
	items []int
	count int
}

// O(1)
func (s *stack) Push(item int) {
	s.items = append(s.items, item)
	s.count++
}

// O(1)
func (s *stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	top := s.items[s.count-1]
	s.items = s.items[:s.count-1]
	s.count--
	return top
}

// O(1)
func (s *stack) IsEmpty() bool {
	return s.count == 0
}

// O(1)
func (s *stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.items[s.count-1]
}

// O(1)
func (s *stack) Size() int {
	return s.count
}

func (s stack) String() string {
	return fmt.Sprintf("%v", s.items)
}

func NewStack() *stack {
	return &stack{items: make([]int, 0)}
}
