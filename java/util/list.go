package util

type List interface {

	//func Size() int

	Size() int
	IsEmpty() bool
	Contains(obj interface{}) bool
	Add(obj interface{}) bool
	Remove(index int) interface{}
	RemoveObj(obj interface{}) bool
	//ContainsAll() bool
	//AddAll(list List) bool
	//RemoveAll(list List) bool
	Clear()
	Get(index int) interface{}
	Set(index int, obj interface{}) interface{}
	IndexOf(obj interface{}) int
}
