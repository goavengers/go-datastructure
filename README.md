# Структуры данных: 10 типов структур данных, которые нужно знать

Структуры данных играют важную роль в процессе разработки ПО, а еще по ним часто задают вопросы на собеседованиях для разработчиков. Хорошая новость в том, что по сути они представляют собой всего лишь специальные форматы для организации и хранения данных.

### Очередь и стек

Начнем с этих двух, т.к. они немного схожи. 
По факту сами они не являются конкретными структурами данных, а объектами реализующий две операции: положить в очередь (push) и взять из очереди (pull). 
Для очереди порядок выдаются объекты в том же порядке — __FIFO__ (First In First Out), для стека в обратном — __LIFO__ (Last In First Out). Ну, есть ещё очередь с приоритетами — это когда первым выдаёт объект с наивысшим приоритетом.

#### Стек

__Стек__ - это базовая структура данных, в которой вы можете только вставлять или удалять элементы в начале стека. 
Он напоминает стопку книг. Если вы хотите взглянуть на книгу в середине стека, вы сначала должны взять книги, лежащие сверху.
Стек считается __LIFO__ (Last In First Out) - это означает, что последний элемент, который добавлен в стек, - это первый элемент, который из него выходит.

Существует три основных операции, которые могут выполняться в стеках: 
- вставка элемента в стек (называемый «push»);
- удаление элемента из стека (называемое «pop»);
- отображение содержимого стека (иногда называемого «pip»).

**Реализация**
<details>
  <summary>Скрытый код</summary>

  ```go
  // some code here
  ```
</details>

**Сложность:**

Алгоритм | В среднем | Худший случай
--- | --- | ---
Space | O(n) | O(n) |
Search | O(n) | O(n) |
Insert | O(1) | O(1) |
Delete | O(1) | O(1) |

#### Очередь

Вы можете думать об этой структуре, как об очереди людей в продуктовом магазине. Стоящий первым будет обслужен первым. Также как очередь.
Если рассматривать очередь с точки доступа к данным, то она является __FIFO__ (First In First Out). 
Это означает, что после добавления нового элемента все элементы, которые были добавлены до этого, должны быть удалены до того, как новый элемент будет удален.

В очереди есть только две основные операции: 
- enqueue;
- dequeue.
 
Enqueue означает вставить элемент в конец очереди, а dequeue означает удаление переднего элемента.

**Реализация**
<details>
  <summary>Скрытый код</summary>

  ```go
  // some code here
  ```
</details>

**Сложность:**

Алгоритм | В среднем | Худший случай
--- | --- | ---
Space | O(n) | O(n) |
Search | O(n) | O(n) |
Insert | O(1) | O(1) |
Delete | O(1) | O(1) |