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

func (a *ArrayList) Contains(obj interface{}) bool {
	return a.IndexOf(obj)>=0
}

func (a *ArrayList) Add(obj interface{}) bool {
	panic("implement me")
}

func (a *ArrayList) Remove(index int) bool {
	panic("implement me")
}

func (a *ArrayList) Clear() {
	panic("implement me")
}

func (a *ArrayList) Get(index int) {
	panic("implement me")
}

func (a *ArrayList) Set(index int, obj interface{}) {
	panic("implement me")
}

func (a *ArrayList) IndexOf(obj interface{})  int{
	panic("implement me")
}
