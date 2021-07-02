package linkedlist

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// List is a doubly-linked list
// The zero value is empty and ready to use.
type List struct {
	head, tail *Node
	len        int
}

// Node is an element within a doubly-linked list
type Node struct {
	next, prev *Node
	Val        interface{}
}

// Next returns the node that follows this one
func (e *Node) Next() *Node {
	return e.next
}

// Prev returns the node that precedes this one
func (e *Node) Prev() *Node {
	return e.prev
}

// NewList creates a new doubly-linked list.
// Elements will be inserted such that the 0th element of the slice
// becomes the head of the list
func NewList(args ...interface{}) *List {
	l := &List{}
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

// First returns the node at the head of the list
func (l *List) First() *Node {
	return l.head
}

// Last returns the node at the tail of the list
func (l *List) Last() *Node {
	return l.tail
}

// PushFront pushes the given element on to the the head of the list
func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	if l.len == 0 {
		l.len, l.head, l.tail = 1, n, n
		return
	}
	n.next = l.head
	l.head.prev = n
	l.head = n
	l.len++
}

// PushBack appends the given element at the tail of the list
func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}
	if l.len == 0 {
		l.len, l.head, l.tail = 1, n, n
		return
	}
	n.prev = l.tail
	l.tail.next = n
	l.tail = n
	l.len++
}

// PopFront removes the element from the head of the list and returns it
func (l *List) PopFront() (interface{}, error) {
	if l.len == 0 {
		return nil, ErrEmptyList
	}
	v := l.head.Val
	l.head = l.head.next
	l.len--
	if l.len == 0 {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return v, nil
}

// PopBack removes the element from the tail of the list and returns it
func (l *List) PopBack() (interface{}, error) {
	if l.len == 0 {
		return nil, ErrEmptyList
	}
	v := l.tail.Val
	l.tail = l.tail.prev
	l.len--
	if l.len == 0 {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return v, nil
}

// Reverse the order of the list
func (l *List) Reverse() *List {
	l.head, l.tail = l.tail, l.head
	curr := l.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.next
	}
	return l
}

// String implements the fmt.Stringer interface
func (l *List) String() string {
	var b strings.Builder
	b.WriteByte('[')
	b.WriteString(strconv.Itoa(l.len))
	b.Write([]byte{']', ' ', '{'})
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Fprint(&b, curr.Val)
		b.WriteByte(',')
	}
	b.WriteByte('}')
	return b.String()
}

// ErrEmptyList is returned if PopFront or PopBack is called on an empty list
var ErrEmptyList = errors.New("cannot pop from an empty list")
