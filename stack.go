package main

import (
	"fmt"
	"sync"
)

func UseStack() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())

	stack.Push("String")
	fmt.Println(stack.Size())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
}

type Stack struct {
	mt      sync.RWMutex
	storage map[int]interface{}
	count   int
}

func NewStack() *Stack {
	return &Stack{
		storage: map[int]interface{}{},
		count:   0,
	}
}

// Adds a value onto the end of the stack
func (s *Stack) Push(value interface{}) {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.storage[s.count] = value
	s.Inc()
}

// Removes and returns the value at the end of the stack
func (s *Stack) Pop() interface{} {
	s.mt.Lock()
	defer s.mt.Unlock()
	if s.Size() == 0 {
		return nil
	}

	s.Dec()
	result := s.storage[s.Size()]
	delete(s.storage, s.Size())
	return result
}

// Returns the value at the end of the stack
func (s *Stack) Peek() interface{} {
	s.mt.RLock()
	defer s.mt.RUnlock()
	return s.storage[s.Size()-1]
}

func (s *Stack) Size() int {
	return s.count
}

func (s *Stack) Inc() {
	s.count++
}

func (s *Stack) Dec() {
	s.count--
}
