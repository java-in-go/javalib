package util

import (
	"strconv"
)

const DEFAULT_CAPACITY = 10

var EMPTY_ELEMENTDATA []*interface{}

type ArrayList struct {
	size        int
	elementData []*interface{}
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
		arrayList.elementData = make([]*interface{}, 0, initialCapacity)
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

func (a *ArrayList) Contains(obj *interface{}) bool {
	return a.IndexOf(obj) >= 0
}

func (a *ArrayList) Add(obj *interface{}) bool {
	a.elementData = append(a.elementData, obj)
	a.size++
	return true
}

func (a *ArrayList) Remove(index int) *interface{} {
	a.rangeCheck(index)
	oldValue := a.elementData[index]
	numMoved := a.size - index - 1
	if numMoved > 1 {

	}
	return oldValue

}

func (a *ArrayList) rangeCheck(index int) {
	if index >= a.size {
		panic("IndexOutOfBoundsException Index: " + strconv.Itoa(index) + ", Size: " + strconv.Itoa(a.size))
	}

}

func (a *ArrayList) Clear() {
	panic("implement me")
}

func (a *ArrayList) Get(index int) {
	panic("implement me")
}

func (a *ArrayList) Set(index int, obj *interface{}) {
	panic("implement me")
}

func (a *ArrayList) IndexOf(obj *interface{}) int {
	panic("implement me")
}
