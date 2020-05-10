package main

import "fmt"

func UseLinkedStack() {
	stack := NewLinkedStack()
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

type LinkedStack struct {
	linkedList LinkedList
}

func NewLinkedStack() LinkedStack {
	return LinkedStack{linkedList: LinkedList{
		count: 0,
		head:  nil,
		tail:  nil,
	}}
}

func (ls *LinkedStack) Peek() interface{} {
	if ls.IsEmpty() {
		return nil
	}

	return ls.linkedList.head.value
}

func (ls *LinkedStack) Push(value interface{}) {
	ls.linkedList.PreAdd(value)
}

func (ls *LinkedStack) Pop() interface{} {
	if ok, removed := ls.linkedList.RemoveHead(); ok {
		return removed.value
	}

	return nil
}

func (ls *LinkedStack) IsEmpty() bool {
	return ls.linkedList.head == nil
}

func (ls *LinkedStack) Size() int {
	return ls.linkedList.Size()
}
