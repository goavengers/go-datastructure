# Структуры данных: 10 типов структур данных, которые необходимо знать каждому

## Содержание

1. [Зачем нужны структуры данных?](#why_are_structures_needed)
2. [Наиболее распространенные структуры данных](#most_popular_structures)
3. [Стек (Stack)](#stack)
    1. [Реализация](#stack_implementation)
    2. [Сложность](#stack_complexity)
    3. [Вопросы о стеке, часто задаваемые на собеседованиях](#stack_interview)
4. [Очереди (Queue)](#queue)
    1. [Реализация](#queue_implementation)
    2. [Сложность](#queue_complexity)
    3. [Вопросы о очередях, часто задаваемые на собеседованиях](#queue_interview)
5. [Множества (Sets)](#sets)
    1. [Реализация](#sets_implementation)
    2. [Сложность](#sets_complexity)
    3. [Вопросы о множествах, часто задаваемые на собеседованиях](#sets_interview)
6. [Связанные списки (Linked Lists)](#linked_lists)
    1. [Реализация](#linked_lists_implementation)
    2. [Сложность](#linked_lists_complexity)
    3. [Реализация стека и очереди на основе связанных списков](#linked_list_stack_queue)
    4. [Двусвязанные списки (Double Linked List)](#linked_list_double)
    5. [Вопросы о связанных списках, часто задаваемые на собеседованиях](#linked_lists_interview)

Структуры данных играют важную роль в процессе разработки ПО, а еще по ним часто задают вопросы на собеседованиях для разработчиков. Хорошая новость в том, что по сути они представляют собой всего лишь специальные форматы для организации и хранения данных.

### <a name="why_are_structures_needed"></a>Зачем нужны структуры данных?

Поскольку структуры данных используются для хранения информации в упорядоченном виде, а данные — самый важный феномен в информатике, истинная ценность структур данных очевидна.

Не важно, какую именно задачу вы решаете, так или иначе вам придется иметь дело с данными, будь то зарплата сотрудника, биржевые котировки, список продуктов для похода в магазин или обычный телефонный справочник.

В зависимости от конкретного сценария, данные нужно хранить в подходящем формате. У нас в распоряжении — ряд структур данных, обеспечивающих нас такими различными форматами.

### <a name="most_popular_structures"></a> Наиболее распространенные структуры данных

Сначала давайте перечислим наиболее распространенные структуры данных, а затем разберем каждую по очереди:

- Массивы
- Динамические массивы
- Множества
- Стеки
- Очереди
- Связные списки
- Деревья
- Графы
- Боры (в сущности, это тоже деревья, но их целесообразно рассмотреть отдельно).
- Хеш-таблицы

### Реализации

Все полные исходники реализаций структур данных вы всегда сможете найти в одноименных фалах расширения `.go`. 
Все структуры предоставляют функцию `Use!Structure!`, где `!Structure!` — это название структуры или псевдоним, однозначно определяющий с какой структурой мы сейчас работаем, например для стека `UseStack`, для очереди `UseQueue`, для структуры стек, основанной на связанных списках псевдоним `UseLinkedStack`.

Полный список таких методов всегда предоставляется в файле `main.go`

### Очередь и стек (Queue & Stack)

Начнем с этих двух, т.к. они немного схожи. 
По факту сами они не являются конкретными структурами данных, а объектами реализующий две операции: положить в очередь (push) и взять из очереди (pull). 
Для очереди порядок выдаются объекты в том же порядке — __FIFO__ (First In First Out), для стека в обратном — __LIFO__ (Last In First Out). Ну, есть ещё очередь с приоритетами — это когда первым выдаёт объект с наивысшим приоритетом.

---

### <a name="stack"></a> Стек

Это коллекция, элементы которой получают по принципу «последний вошел, первый вышел» (Last-In-First-Out или __LIFO__). 
Это значит, что мы будем иметь доступ только к последнему добавленному элементу.

В отличие от списков, мы не можем получить доступ к произвольному элементу стека. 
Мы можем только добавлять или удалять элементы с помощью специальных методов. 
У стека нет также метода Contains, как у списков. 
Кроме того, у стека нет итератора. 
Для того, чтобы понимать, почему на стек накладываются такие ограничения, давайте посмотрим на то, как он работает и как используется.

Наиболее часто встречающаяся аналогия для объяснения стека — стопка тарелок. 
Вне зависимости от того, сколько тарелок в стопке, мы всегда можем снять верхнюю. 
Чистые тарелки точно так же кладутся на верх стопки, и мы всегда будем первой брать ту тарелку, которая была положена последней.

![stack](img/stack.jpg)

Если мы положим, например, красную тарелку, затем синюю, а затем зеленую, то сначала надо будет снять зеленую, потом синюю, и, наконец, красную. 
Главное, что надо запомнить — тарелки всегда ставятся и на верх стопки. 
Когда кто-то берет тарелку, он также снимает ее сверху. 
Получается, что тарелки разбираются в порядке, обратном тому, в котором ставились.

Существует три основных операции, которые могут выполняться в стеках: 
- вставка элемента в стек (называемый «push»);
- удаление элемента из стека (называемое «pop»);
- отображение содержимого стека (иногда называемого «peak»).

Теперь, когда мы понимаем, как работает стек, введем несколько терминов. 
Операция добавления элемента на стек называется «push», удаления — «pop». 
Последний добавленный элемент называется верхушкой стека, или «top», и его можно посмотреть с помощью операции «peek». 

<a name="stack_implementation"></a> **Реализация**

<details>
  <summary>Скрытый код</summary>

  ```go
  type Stack struct {
  	mt         sync.RWMutex
  	collection map[int]interface{}
  	count      int
  }
  
  func NewStack() *Stack {
  	return &Stack{
  		collection: map[int]interface{}{},
  		count:      0,
  	}
  }
  
  // Adds a value onto the end of the stack
  // Добавляет элемент на вершину стека.
  // Сложность: O(1).
  func (s *Stack) Push(value interface{}) {
  	s.mt.Lock()
  	defer s.mt.Unlock()
  	s.collection[s.count] = value
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
  	result := s.collection[s.Size()]
  	delete(s.collection, s.Size())
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
  
  	return s.collection[s.Size()-1]
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
  ```
</details>

<a name="stack_usage"></a> **Примнение**

<details>
  <summary>Обра́тная по́льская запись (англ. Reverse Polish notation, RPN)</summary>

  ```go
  // todo coming soon
  ```
</details>

<a name="stack_complexity"></a> **Сложность:**

Алгоритм | В среднем | Худший случай
--- | --- | ---
Space | O(n) | O(n) |
Search | O(n) | O(n) |
Insert | O(1) | O(1) |
Delete | O(1) | O(1) |

<a name="stack_interview"></a> **Вопросы о стеке, часто задаваемые на собеседованиях:**

- Вычислить постфиксное выражение при помощи стека
- Отсортировать значения в стеке
- Проверить сбалансированные скобки в выражении

---

### Очередь

Вы можете думать об этой структуре, как об очереди людей в продуктовом магазине. Стоящий первым будет обслужен первым. Также как очередь.
Если рассматривать очередь с точки доступа к данным, то она является __FIFO__ (First In First Out). 
Это означает, что после добавления нового элемента все элементы, которые были добавлены до этого, должны быть удалены до того, как новый элемент будет удален.

В очереди есть только две основные операции и две вспомогательные: 
- enqueue;
- dequeue;
- peek;
- count.
 
Enqueue означает вставить элемент в конец очереди, а dequeue означает удаление переднего элемента.

**Реализация**
<details>
  <summary>Скрытый код</summary>

  ```go
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
  ```
</details>

**Сложность:**

Алгоритм | В среднем | Худший случай
--- | --- | ---
Space | O(n) | O(n) |
Search | O(n) | O(n) |
Insert | O(1) | O(1) |
Delete | O(1) | O(1) |

**Вопросы о стеке, часто задаваемые на собеседованиях:**

- Реализуйте стек при помощи очереди
- Обратите первые k элементов в очереди
- Сгенерируйте двоичные числа от 1 до n при помощи очереди

---

### Множества (Sets)

Это коллекция, которая реализует основные математические операции над множествами: 

__Пересечения__ (intersection) - если заданы два множества, эта функция вернет другое множество, содержащее элементы, которые имеются и в первом и во втором множестве;

```go
[1, 2, 3, 4] intersect [3, 4, 5, 6] = [3, 4]
```

__Объединение__ (union) - объединяет все элементы из двух разных множеств и возвращает результат, как новый набор (без дубликатов);

```go
[1, 2, 3, 4] union [3, 4, 5, 6] = [1, 2, 3, 4, 5, 6]
```

__Разность__ (difference) - вернет список элементов, которые находятся в одном множестве, но НЕ повторяются в другом;

```go
[1, 2, 3, 4] difference [3, 4, 5, 6] = [1, 2]
```

__Симметрическая разность__ (symmetric difference) - это все элементы, которые содержатся только в одном из рассматриваемых множеств;

```go
[1, 2, 3, 4] symmetric difference [3, 4, 5, 6] = [1, 2, 5, 6]
```

Вы, возможно, заметили, что симметрическая разность — это «пересечение наоборот». Учитывая это, давайте попробуем найти ее, используя уже имеющиеся операции.

Итак, мы хотим получить множество, которое содержит все элементы из двух множеств, за исключением тех, которые есть в обоих. Другими словами, мы хотим получить разность объединения двух множеств и их пересечения.

Или, если рассматривать по шагам:

```go
[1, 2, 3, 4] union [3, 4, 5, 6] = [1, 2, 3, 4, 5, 6]

[1, 2, 3, 4] intersection [3, 4, 5, 6] = [3, 4]

[1, 2, 3, 4, 5, 6] set difference [3, 4] = [1, 2, 5, 6]
```

Что дает нам нужный результат:

```go
[1, 2, 5, 6]
```

__Подмножество__ (subset) - возвращает булево значение, показывающее, содержит ли одно множество все элементы другого множества. 

```go
[1, 2, 3] is subset [0, 1, 2, 3, 4, 5] = true
```

В то время как:

```go
[1, 2, 3] is subset [0, 1, 2] = false
```

Эту проверку можно провести, используя имещиеся методы: difference и intersection, пустой результат при разности и множество с таким же количеством элементов при пересечении говорит о том, что все элементы первого множества содержатся во втором множестве;

```go
[1, 2, 3] difference [0, 1, 2, 3, 4, 5] = [] = true

[1, 2, 3] intersection [0, 1, 2, 3, 4, 5] = [1, 2, 3] = true
```

В общем случае, конечно, множества (sets) может иметь метод __subset__ (который может быть реализован более эффективно). Однако стоит помнить, что это уже не какая-то новая возможность, а просто другое применение существующих.

**Помните:** Множества хранят данные без определенного порядка и без повторяющихся значений.

**Реализация**

<details>
  <summary>Скрытый код</summary>
  
  ```go
  type Set struct {
  	collection []interface{}
  }
  
  func (s *Set) Collections() []interface{} {
  	return s.collection
  }
  
  // Поведение: Добавляет элементы в множество.
  // Если элемент уже присутствует в множестве, возвращается false, иначе true.
  // Сложность: O(n)
  func (s *Set) Add(value interface{}) bool {
  	if _, exist := s.Contains(value); exist == false {
  		s.collection = append(s.collection, value)
  		return true
  	}
  
  	return false
  }
  
  // Поведение: Возвращает index, true, если множество содержит указанный элемент.
  // В противном случае возвращает index, false.
  // Сложность: O(n)
  func (s *Set) Contains(value interface{}) (int, bool) {
  	for k, v := range s.collection {
  		if v == value {
  			return k, true
  		}
  	}
  
  	return 0, false
  }
  
  // Поведение: Удаляет указанный элемент из множества и возвращает true.
  // В случае, если элемент нет и он не удален, возвращает false.
  // Сложность: O(n)
  func (s *Set) Remove(value interface{}) bool {
  	if i, exist := s.Contains(value); exist == true {
  		length := s.Size()
  
  		s.collection[i] = s.collection[length-1]
  		s.collection = s.collection[:length-1]
  
  		return true
  	}
  
  	return false
  }
  
  func (s *Set) Size() int {
  	return len(s.collection)
  }
  
  // Поведение: Возвращает множество, полученное операцией объединения его с указанным.
  // Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
  func (s *Set) Union(set *Set) *Set {
  	union := &Set{}
  
  	for _, v := range s.Collections() {
  		union.Add(v)
  	}
  
  	for _, v := range set.Collections() {
  		union.Add(v)
  	}
  
  	return union
  }
  
  // Поведение: Возвращает множество, полученное операцией пересечения его с указанным.
  // Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
  func (s *Set) Intersection(set *Set) *Set {
  	intersection := &Set{}
  
  	for _, v := range s.Collections() {
  		if _, exist := set.Contains(v); exist == true {
  			intersection.Add(v)
  		}
  	}
  
  	return intersection
  }
  
  // Поведение: Возвращает множество, являющееся разностью текущего с указанным.
  // Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
  func (s *Set) Difference(set *Set) *Set {
  	difference := &Set{}
  
  	for _, v := range s.Collections() {
  		if _, exist := set.Contains(v); exist == false {
  			difference.Add(v)
  		}
  	}
  
  	return difference
  }
  
  // Поведение: Возвращает true, если второе множество является подмножеством первого, в противном случае возвращает false.
  // Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
  func (s *Set) Subset(set *Set) bool {
  	if s.Size() > set.Size() {
  		return false
  	}
  
  	return s.Difference(set).Size() == 0
  }
  
  // Поведение: Возвращает множество, являющееся симметрической разностью текущего с указанным.
  // Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
  func (s *Set) SymmetricDifference(set *Set) *Set {
  	a := s.Difference(set)
  	b := set.Difference(s)
  
  	return a.Union(b)
  }
  ```
</details>

---

### Связанные списки (Linked List)

Основное назначение связного списка — предоставление механизма для хранения и доступа к произвольному количеству данных. Как следует из названия, это достигается связыванием данных вместе в список.

Для начала определим несколько терминов, чтобы в дальнейшем не возникло недопонимания:

- __указатель__ содержит адрес участка памяти, хранящего определённые данные (для указателей также допустимо нулевое значение);
- __ссылка__, в отличие от указателя, всегда должна указывать на определённый адрес;
- __структура данных__ — способ группировки информации, который может быть реализован на любом языке программирования.

Простейший способ создать односвязный список — поочерёдно создать и связать узлы:

```go
type Node struct {
    value interface{}
    next *Node
}

first := &Node{3}
moddle := &Node{5}
last := &Node{7}

zero.next = moddle

// +-----+------+    +-----+------+
// |  3  |  *---+--->|  5  | null +
// +-----+------+    +-----+------+

moddle.next = last

// +-----+------+    +-----+------+   +-----+------+
// |  3  |  *---+--->|  5  |  *---+-->|  7  | null +
// +-----+------+    +-----+------+   +-----+------+

```

**Реализация**

<details>
  <summary>Скрытый код</summary>
  
  ```go
  type LinkedList struct {
  	count int
  	head *Node
  	tail *Node
  }
  
  type Node struct {
  	value interface{}
  	next *Node
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
  ```
</details>

<a name="linked_list_stack_queue"></a> **Реализация стека и очереди на основе связанных списков**

С помощью связанных списков можно можно создавать такие структуры данных, как [Стеки (Stack)](#stack) и [Очереди (Queue)](#queue), которые мы уже реализовывали выше, давайте попробуем реализовать их с помощью связанных списков.

Для этого, нам понадобится немного модифицировать существующий код связанных списков и добавить к ним несколько методов, которые сильно напоминают методы из двухсвязанных списков [(Double Linked List)](#linked_list_double) о которых мы поговорим чуть позже.

```go
// данный метод добавляет элемент перед первым элементом в списке
func (l *LinkedList) PreAdd(value interface{}) {
	newNode := &Node{
		value: value,
		next:  l.head,
	}

	l.head = newNode

	if l.tail == nil {
		l.tail = newNode
	}

	l.count++
}

// следующие два метода удаляют с начала и конца соответственно названиям метода
func (l *LinkedList) RemoveHead() (bool, *Node) {
	if l.head == nil {
		return false, nil
	}

	tmpHead := l.head

	if l.head.next != nil {
		l.head = l.head.next
	} else {
		l.head = nil
		l.head = nil
	}

	l.count--
	return true, tmpHead
}

func (l *LinkedList) RemoveTail() (bool, *Node) {
	tmpTail := l.tail

	if l.head.next == nil {
		// There is only one node in linked list.

		l.head = nil
		l.tail = nil
		l.count--
		return true, tmpTail
	}

	// If there are many nodes in linked list...
	// Rewind to the last node and delete "next" link for the node before the last one.

	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next.next == nil {
			currentNode.next = nil
		} else {
			currentNode = currentNode.next
		}
	}

	l.count--
	l.tail = currentNode
	return true, tmpTail
}
```

Теперь, когда мы добавили необходимые методы, мы можем реализовывать структуру данных Стек с помощью связанных списков:

<details>
  <summary>Реализация структуры __Стек__ с помощью связанного списка</summary>
  
  ```go
  type LinkedStack struct {
  	linkedList LinkedList
  }
  
  func NewLinkedStack() LinkedStack {
  	return LinkedStack{linkedList: LinkedList {
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
  ```
</details>

Таким же образом реализуем структуру очередь:

<details>
  <summary>Реализация структуры __Очередь__ с помощью связанного списка</summary>
  
  ```go
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
  ```
</details>

Вот так вот легко реализовывать структуры данных поверх другой структуры данных. 
Сравните полученный код с "нативной" реализацией и вы увидите, что код стал короче и компактнее, а так же то, что код выглядит более унифицированным.

**Вопросы о стеке, часто задаваемые на собеседованиях:**

- Обратите связный список
- Найдите петлю в связном списке
- Возвратите N-ный узел с начала связного списка
- Удалите из связного списка дублирующиеся значения
