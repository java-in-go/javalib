package util

import (
	"container/list"
	"reflect"
	"strconv"
)

type LinkedList struct {
	list *list.List
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		list: list.New(),
		size: 0,
	}
}

func (a *LinkedList) Size() int {
	return a.size
}

func (a *LinkedList) IsEmpty() bool {
	return a.size == 0
}

func (a *LinkedList) Contains(obj interface{}) bool {
	return a.IndexOf(obj) > -1
}

func (a *LinkedList) Add(obj interface{}) bool {
	a.list.PushBack(obj)
	a.size++
	return true
}

func (a *LinkedList) Remove(index int) interface{} {
	a.checkElementIndex(index)
	value := a.list.Remove(a.node(index))
	a.size--
	return value
}

func (a *LinkedList) node(index int) *list.Element {
	// index 在前半截
	if index < (a.size >> 1) {
		x := a.list.Front()
		for i := 0; i < index; i++ {
			x = x.Next()
		}
		return x
	} else {
		x := a.list.Back()
		for i := a.size - 1; i > index; i-- {
			x = x.Prev()
		}
		return x
	}
}
func (a *LinkedList) RemoveObj(obj interface{}) bool {
	if obj == nil {
		for x := a.list.Front(); x != nil; x = x.Next() {
			if x.Value == nil {
				a.list.Remove(x)
				a.size--
				return true
			}
		}
	} else {
		for x := a.list.Front(); x != nil; x = x.Next() {
			if reflect.DeepEqual(x.Value, obj) {
				a.list.Remove(x)
				a.size--
				return true
			}
		}
	}
	return false
}

func (a *LinkedList) checkElementIndex(index int) {
	if !a.isElementIndex(index) {
		panic("IndexOutOfBoundsException Index: " + strconv.Itoa(index) + ", Size: " + strconv.Itoa(a.size))
	}
}
func (a *LinkedList) isElementIndex(index int) bool {
	return a.size >= 0 && index < a.size
}

func (a *LinkedList) Clear() {
	for x := a.list.Front(); x != nil; {
		next := x.Next()
		a.list.Remove(x)
		x = next
	}
	a.size = 0
}

func (a *LinkedList) Get(index int) interface{} {
	a.checkElementIndex(index)
	return a.node(index).Value
}

func (a *LinkedList) Set(index int, obj interface{}) interface{} {
	a.checkElementIndex(index)
	node := a.node(index)
	oldValue := node.Value
	node.Value = obj
	return oldValue
}

func (a *LinkedList) IndexOf(obj interface{}) int {
	index := 0
	if obj == nil {
		for x := a.list.Front(); x != nil; x = x.Next() {
			if x.Value == nil {
				return index
			}
			index++
		}
	} else {
		for x := a.list.Front(); x != nil; x = x.Next() {
			if reflect.DeepEqual(x.Value, obj) {
				return index
			}
			index++
		}
	}
	return -1
}

func (a *LinkedList) ForEach(f func(value interface{})) {
	l := a.list
	for e := l.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}
func (a *LinkedList) ForEachIndex(f func(index int, value interface{})) {
	index := 0
	for e := a.list.Front(); e != nil; e = e.Next() {
		f(index, e.Value)
		index++
	}
}

func (a *LinkedList) Map(f func(value interface{}) interface{}) List {
	linkedList := NewLinkedList()
	l := a.list
	for e := l.Front(); e != nil; e = e.Next() {
		linkedList.Add(f(e.Value))
	}
	return linkedList
}

func (a *LinkedList) Reduce(f func(v1 interface{}, v2 interface{}) interface{}) interface{} {
	foundAny := false
	var result interface{}
	l := a.list
	for e := l.Front(); e != nil; e = e.Next() {
		if !foundAny {
			foundAny = true
			result = e.Value
			continue
		}
		result = f(result, e.Value)
	}
	if foundAny {
		return result
	}
	return nil
}
