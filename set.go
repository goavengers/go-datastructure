package main

import "fmt"

func UseSet() {
	setA := &Set{}
	setB := &Set{}

	setA.Add(10)
	setA.Add(20)

	setB.Add(10)
	setB.Add(20)
	setB.Add(50)
	setB.Add(60)
	setB.Add(70)

	fmt.Println(setA.Subset(setB))
	fmt.Println(setA.Intersection(setB).Collections())
	fmt.Println(setB.Difference(setA).Collections())

	fmt.Println(setA.Collections())
	setA.Remove(10)
	setA.Add(20)
	fmt.Println(setA.Collections())

	fmt.Println(setA.Contains(10))

	setA.Add(10)
	setA.Add(60)
	fmt.Println(setA.Collections())
	fmt.Println(setB.Collections())
	fmt.Println(setB.SymmetricDifference(setA))
}

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
