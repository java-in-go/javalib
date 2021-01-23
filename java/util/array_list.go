package util

import (
	"reflect"
	"strconv"
)

const DefaultCapacity = 10

var EmptyElementData []interface{}

type ArrayList struct {
	size        int
	elementData []interface{}
}

func NewArrayList() *ArrayList {
	return NewArrayListCap(DefaultCapacity)
}
func NewArrayListCap(initialCapacity int) *ArrayList {
	arrayList := &ArrayList{}
	if initialCapacity > 0 {
		arrayList.elementData = make([]interface{}, 0, initialCapacity)
	} else if initialCapacity == 0 {
		arrayList.elementData = EmptyElementData
	} else {
		panic("Illegal Capacity: " + strconv.Itoa(initialCapacity))
	}
	return arrayList
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList) Contains(obj interface{}) bool {
	return a.IndexOf(obj) >= 0
}

func (a *ArrayList) Add(obj interface{}) bool {
	a.elementData = append(a.elementData, obj)
	a.size++
	return true
}
func (a *ArrayList) initCapacity() {
	if cap(a.elementData) == 0 {
		a.elementData = make([]interface{}, 0, DefaultCapacity)
	}
}
func (a *ArrayList) ensureCapacityInternal(index int) int {
	if len(a.elementData) == 0 {
		return MaxInt(DefaultCapacity, index)
	}
	return index
}
func (a *ArrayList) ensureExplicitCapacity(minCapacity int) {
	oldCapacity := cap(a.elementData)
	if minCapacity-oldCapacity > 0 {
		newCapacity := oldCapacity + (oldCapacity >> 1)
		if newCapacity-minCapacity < 0 {
			newCapacity = minCapacity
		}
		newElementData := make([]interface{}, 0, newCapacity)
		copy(a.elementData, newElementData)
		a.elementData = newElementData
	}
}

func (a *ArrayList) Remove(index int) interface{} {
	a.rangeCheck(index)
	oldValue := a.elementData[index]
	numMoved := a.size - index - 1
	if numMoved > 0 {
		a.elementData = append(a.elementData[:0], a.elementData[index+1])
	}
	a.size--
	return oldValue
}

func (a *ArrayList) RemoveObj(obj interface{}) bool {
	index := a.IndexOf(obj)
	if index > -1 {
		a.Remove(index)
		return true
	}
	return false
}

func (a *ArrayList) rangeCheck(index int) {
	if index >= a.size {
		panic("IndexOutOfBoundsException Index: " + strconv.Itoa(index) + ", Size: " + strconv.Itoa(a.size))
	}

}

func (a *ArrayList) Clear() {
	for i := range a.elementData {
		a.elementData[i] = nil
	}
	a.size = 0
}

func (a *ArrayList) Get(index int) interface{} {
	a.rangeCheck(index)
	return a.elementData[index]
}

//  Replaces the element at the specified position in this list with
//  the specified element.
//
//  index index of the element to replace
//  element element to be stored at the specified position
//  @returns the element previously at the specified position
func (a *ArrayList) Set(index int, obj interface{}) interface{} {
	a.rangeCheck(index)
	oldValue := a.elementData[index]
	a.elementData[index] = obj
	return oldValue
}

// Returns the index of the first occurrence of the specified element in this list,
//or -1 if this list does not contain the element.
func (a *ArrayList) IndexOf(obj interface{}) int {
	if obj == nil {
		for i := 0; i < a.size; i++ {
			if a.elementData[i] == nil {
				return i
			}
		}
	} else {
		for i := 0; i < a.size; i++ {
			datum := a.elementData[i]
			if datum == obj || reflect.DeepEqual(datum, obj) {
				return i
			}
		}
	}
	return -1
}

func (a *ArrayList) ForEach(f func(value interface{})) {
	for i := 0; i < a.size; i++ {
		f(a.Get(i))
	}
}

func (a *ArrayList) ForEachIndex(f func(index int, value interface{})) {
	for i := 0; i < a.size; i++ {
		f(i, a.Get(i))
	}
}
func (a *ArrayList) Map(f func(value interface{}) interface{}) List {
	newList := NewArrayListCap(a.size)
	for i := 0; i < a.size; i++ {
		newList.Add(f(a.Get(i)))
	}
	return newList
}

func (a *ArrayList) Reduce(f func(v1 interface{}, v2 interface{}) interface{}) interface{} {
	foundAny := false
	var result interface{}
	for i := 0; i < a.size; i++ {
		v := a.Get(i)
		if !foundAny {
			foundAny = true
			result = v
		}
		result = f(result, v)
	}
	if foundAny {
		return result
	}
	return nil
}
