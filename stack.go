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
// Добавляет элемент на вершину стека.
// Сложность: O(1).
func (s *Stack) Push(value interface{}) {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.storage[s.count] = value
	s.Inc()
}

// Removes and returns the value at the end of the stack
// Удаляет элемент с вершины стека и возвращает его. Если стек пустой, возвращает nil
// Т.к.`Push` добавляет элементы в конец списка, поэтому забирать их будет также с конца.
// Сложность: O(1).
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
// Возвращает верхний элемент стека, но не удаляет его.
// Сложность: O(1).
func (s *Stack) Peek() interface{} {
	s.mt.RLock()
	defer s.mt.RUnlock()

	if s.Size() == 0 {
		return nil
	}

	return s.storage[s.Size()-1]
}

// Get count elements in stack
// Возвращает количество элементов в стеке.
// Зачем нам знать, сколько элементов находится в стеке, если мы все равно не имеем к ним доступа?
// С помощью этого поля мы можем проверить, есть ли элементы на стеке или он пуст.
// Сложность: O(1).
func (s *Stack) Size() int {
	return s.count
}

func (s *Stack) Inc() {
	s.count++
}

func (s *Stack) Dec() {
	s.count--
}
