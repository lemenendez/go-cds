package stacks

import (
	"errors"
	"fmt"
)

var ErrEmpty = errors.New("data struct is empty")

// Node a Stack
type Node[T any] struct {
	Next *Node[T]
	Val  T
}

// String() converts stack into a string
func (n Node[any]) String() string {
	return fmt.Sprintf("%v", n.Val)
}

// NewNode create new Node
func NewNode[T any](v T) *Node[any] {
	n := &Node[any]{}
	n.Val = v
	return n
}

// Stack Linked List Stack implementation
type Stack[T any] struct {
	size int
	head *Node[T]
}

func NewStack[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	// dummy node
	s.head = &Node[T]{}
	return s
}

func (s Stack[T]) toSlice() []T {
	a := make([]T, s.size)
	counter := 0
	curr := s.head.Next
	for curr != nil {
		a[counter] = curr.Val
		curr = curr.Next
		counter++
	}
	return a
}

func (s Stack[any]) String() string {
	return fmt.Sprintf("%v", s.toSlice())
}

func (s *Stack[T]) Push(val T) {
	nn := &Node[T]{}
	nn.Val = val
	nn.Next = s.head.Next
	s.head.Next = nn
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		return *new(T), ErrEmpty
	}
	n := s.head.Next
	s.head.Next = s.head.Next.Next
	s.size--
	return n.Val, nil
}

func (s Stack[any]) Size() int {
	return s.size
}
