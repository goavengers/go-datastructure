package main

import (
	"fmt"
	"sync"
)

func UseQueue() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Collections())
	queue.Dequeue()
	fmt.Println(queue.Peek())
	fmt.Println(queue.Collections())
}

type Queue struct {
	mt         sync.RWMutex
	collection []interface{}
}

// Добавляет элемент в очередь.
// Новые элементы очереди можно добавлять как в начало списка, так и в конец.
// Важно только, чтобы элементы доставались с противоположного края.
// Сложность: O(1).
func (q *Queue) Enqueue(value interface{}) {
	q.mt.Lock()
	defer q.mt.Unlock()
	q.collection = append(q.collection, value)
}

// Удаляет первый помещенный элемент из очереди и возвращает его.
// Если очередь пустая, возвращает nil
// Поскольку мы вставляем элементы в конец списка, убирать мы их будем с начала.
// Сложность: O(1).
func (q *Queue) Dequeue() interface{} {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.isEmpty() {
		return nil
	}

	val := q.collection[0]
	q.collection = q.collection[1:]

	return val
}

// Возвращает элемент, который вернет следующий вызов метода Dequeue.
// Очередь остается без изменений. Если очередь пустая, возврщается nil
// Сложность: O(1)
func (q *Queue) Peek() interface{} {
	q.mt.RLock()
	defer q.mt.RUnlock()

	if q.isEmpty() {
		return nil
	}

	return q.collection[0]
}

// Возвращает количество элементов в очереди или 0, если очередь пустая.
// Сложность: O(1).
func (q *Queue) Count() int {
	return len(q.collection)
}

func (q *Queue) isEmpty() bool {
	return q.Count() == 0
}

func (q *Queue) Collections() []interface{} {
	return q.collection
}
