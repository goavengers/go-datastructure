package main

import "fmt"

func UseLinkedQueue() {
	queue := &LinkedQueue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.ToString()
	queue.Dequeue()
	fmt.Println(queue.Peek())
	queue.ToString()
}

type LinkedQueue struct {
	linkedList LinkedList
}

func (lq *LinkedQueue) Enqueue(value interface{}) {
	lq.linkedList.Add(value)
}

func (lq *LinkedQueue) Dequeue() interface{} {
	if ok, removed := lq.linkedList.RemoveHead(); ok {
		return removed.value
	}

	return nil
}

func (lq *LinkedQueue) Peek() interface{} {
	if lq.linkedList.head == nil {
		return nil
	}

	return lq.linkedList.head.value
}

func (lq *LinkedQueue) ToString() {
	lq.linkedList.PrintNodes()
}
