package main

import "fmt"

func UseLinkedList() {
	linkedList := &LinkedList{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.PrintNodes()

	if index, removed := linkedList.Remove(3); removed {
		fmt.Println(fmt.Sprintf("Removed index %d", index))
		linkedList.PrintNodes()
	}

	if index, exist := linkedList.Contains(2); exist {
		fmt.Println(fmt.Sprintf("Exist index %d", index))
	}
}

type LinkedList struct {
	count int
	head  *Node
	tail  *Node
}

type Node struct {
	value interface{}
	next  *Node
}

func (l *LinkedList) Add(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.count++
}

func (l *LinkedList) Remove(element interface{}) (int, bool) {
	current := l.head
	var previous *Node

	index := 0
	for current != nil {
		// find element
		if current.value == element {

			// is not first node
			if previous != nil {

				// before: Head -> 1 -> 2 -> 3 --> null
				// after:  Head -> 1 -> 2 -------> null
				previous.next = current.next

				// set last node if current already is last
				if current.next == nil {
					l.tail = previous
				}
			} else {
				// before: Head -> 3 -> 5
				// after:  Head ------> 5

				// Head -> 3 -> null
				// Head ------> null

				l.head = current.next
				// check is empty list
				if l.head == nil {
					l.tail = nil
				}
			}

			l.count--
			return index, true
		}

		index++
		previous = current
		current = current.next
	}

	return index, false
}

func (l *LinkedList) Contains(element interface{}) (int, bool) {
	current := l.head

	index := 0
	for current != nil {
		if current.value == element {
			return index, true
		}

		current = current.next
		index++
	}

	return index, false
}

func (l *LinkedList) Size() int {
	return l.count
}

// help method
func (l *LinkedList) PrintNodes() {
	currentNode := l.head
	chain := ""

	for currentNode.next != nil {
		chain += fmt.Sprintf("| %v : next | --> ", currentNode.value)
		currentNode = currentNode.next
	}

	// last node
	chain += fmt.Sprintf("| %v : next |", currentNode.value)
	fmt.Println(chain)
}
