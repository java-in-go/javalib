package util

import "container/list"

type LinkedList struct {
	list list.List
}

func (a *LinkedList) Size() int {
	return a.list.Len()
}

func (a *LinkedList) IsEmpty() bool {
	return a.list.Len() == 0
}

func (a *LinkedList) Contains(obj interface{}) bool {
	return a.IndexOf(obj) > -1
}

func (a *LinkedList) Add(obj interface{}) bool {
	a.list.PushBack(obj)
	return true
}

func (a *LinkedList) Remove(index int) interface{} {
	//a.list.Remove()
	panic("implement me")
}

func (a *LinkedList) RemoveObj(obj interface{}) bool {
	if obj == nil {

	}
	panic("implement me")
}

func checkElementIndex(index int) {

}

func (a *LinkedList) Clear() {
	panic("implement me")
}

func (a *LinkedList) Get(index int) interface{} {
	panic("implement me")
}

func (a *LinkedList) Set(index int, obj interface{}) interface{} {
	panic("implement me")
}

func (a *LinkedList) IndexOf(obj interface{}) int {
	panic("implement me")
}
