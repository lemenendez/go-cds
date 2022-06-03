package stacks_test

import (
	"github.com/lemenendez/go-cds/stacks"
	"testing"
)

func TestNewNode2(t *testing.T) {
	n := stacks.NewNode[int](89)
	t.Log(n)
	another := stacks.NewNode[string]("my str")
	t.Log(another)
}

func TestStack(t *testing.T) {
	s := stacks.NewStack[int]()
	s.Push(99)
	t.Log(s)
	s.Push(88)
	t.Log(s)
	s.Push(70)
	t.Log(s)
}
