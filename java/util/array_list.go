package util

import (
	"reflect"
	"strconv"
)

const DEFAULT_CAPACITY = 10

var EMPTY_ELEMENTDATA []interface{}

type ArrayList struct {
	size        int
	elementData []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		size:        0,
		elementData: EMPTY_ELEMENTDATA,
	}
}
func NewArrayListCap(initialCapacity int) *ArrayList {
	arrayList := &ArrayList{}
	if initialCapacity > 0 {
		arrayList.elementData = make([]interface{}, 0, initialCapacity)
	} else if initialCapacity == 0 {
		arrayList.elementData = EMPTY_ELEMENTDATA
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
	a.ensureCapacityInternal(a.size + 1)
	a.elementData = append(a.elementData, obj)
	a.size++
	return true
}
func (a *ArrayList) ensureCapacityInternal(index int) int {
	elementData := &a.elementData
	emptyElementData := &EMPTY_ELEMENTDATA
	if elementData == emptyElementData {
		return MaxInt(DEFAULT_CAPACITY, index)
	}
	return index
}
func (a *ArrayList) ensureExplicitCapacity(minCapacity int) {
	if minCapacity-len(a.elementData) > 0 {
		a.elementData = make([]interface{}, 0, minCapacity)
	}
}
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (a *ArrayList) Remove(index int) interface{} {
	a.rangeCheck(index)
	oldValue := a.elementData[index]
	numMoved := a.size - index - 1
	if numMoved > 1 {
		a.elementData = append(a.elementData[:0], a.elementData[index+1])
	}
	a.size--
	return oldValue

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
//  @param index index of the element to replace
//  @param element element to be stored at the specified position
//  @returns the element previously at the specified position
func (a *ArrayList) Set(index int, obj interface{}) interface{} {
	a.rangeCheck(index)
	oldValue := a.elementData[index]
	a.elementData[index] = obj
	return oldValue
}

// Returns the index of the first occurrence of the specified element in this list,
//or -1 if this list does not contain the element.
//More formally, returns the lowest index i such that (o==null ? get(i)==null : o.equals(get(i))),
//or -1 if there is no such index.
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
