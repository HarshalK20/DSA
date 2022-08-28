package linkedlist

import (
	"strconv"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedlist struct {
	head  *node
	tail  *node
	count int
}

// O(n)
func (l *linkedlist) InsertAt(item int, index int) {
	if index < 0 || index >= l.count {
		return
	}

	if index == 0 {
		l.InsertFirst(item)
		return
	}

	for i, j := l.head, 0; i != nil; i, j = i.next, j+1 {
		if j+1 == index {
			n := node{value: item, next: i.next}
			i.next = &n
			l.count++
			break
		}
	}
}

// O(1)
func (l *linkedlist) InsertFirst(item int) {
	n := node{value: item}

	if l.count == 0 {
		l.head = &n
		l.tail = &n
	} else {
		n.next = l.head
		l.head = &n
	}

	l.count++
}

// O(n)
func (l *linkedlist) Insert(item int) {
	n := node{value: item, next: nil}

	if l.count == 0 {
		l.head = &n
		l.tail = &n
	} else {
		i := l.head
		for i.next != nil {
			i = i.next
		}

		i.next = &n
		l.tail = &n
	}

	l.count++
}

// O(n)
func (l *linkedlist) DeleteAt(index int) int {
	if index < 0 || index >= l.count {
		return -1
	}

	if index == 0 {
		return l.DeleteFirst()
	}

	var last int
	for i, j := l.head, 0; i != nil; i, j = i.next, j+1 {
		if j+1 == index {
			if index == l.count-1 {
				i.next = nil
				l.tail = i
			} else {
				last = i.next.value
				i.next = i.next.next
			}
			l.count--
			break
		}
	}
	return last
}

// O(1)
func (l *linkedlist) DeleteFirst() int {
	if l.count == 0 {
		return -1
	}

	var last int
	if l.count == 1 {
		last = l.head.value
		l.head = nil
		l.tail = nil
	} else {
		last = l.head.value
		l.head = l.head.next
	}
	l.count--
	return last
}

// O(n)
func (l *linkedlist) DeleteLast() int {

	var last int
	if l.count == 0 {
		return -1
	} else if l.count == 1 {
		last = l.head.value
		l.head = nil
		l.tail = nil
	} else {
		i, j := l.head, l.head.next
		for j.next != nil {
			i = i.next
			j = j.next
		}

		last = j.value
		i.next = nil
		l.tail = i
	}
	l.count--
	return last
}

// O(n)
func (l *linkedlist) Delete(item int) {

	if l.head.value == item {
		l.DeleteFirst()
		return
	}

	i := l.head
	for i.next != nil {
		if i.next.value == item {
			i.next = i.next.next
			l.count--
		}
		i = i.next
	}
}

// O(n)
func (l *linkedlist) Contains(item int) bool {
	return l.IndexOf(item) != -1
}

// O(n)
func (l *linkedlist) IndexOf(item int) int {
	for i, j := l.head, 0; i != nil; i, j = i.next, j+1 {
		if i.value == item {
			return j
		}
	}

	return -1
}

func (l *linkedlist) String() string {
	s := make([]string, 0)
	for i := l.head; i != nil; i = i.next {
		s = append(s, strconv.Itoa(i.value))
	}
	return strings.Join(s, " -> ")
}

func NewLinkedList() *linkedlist {
	return &linkedlist{}
}
