package util

type List interface {

	//func Size() int

	Size() int
	IsEmpty() bool
	Contains(obj interface{}) bool
	Add(obj interface{}) bool
	Remove(index int) bool
	//ContainsAll() bool
	//AddAll(list List) bool
	//RemoveAll(list List) bool
	Clear()
	Get(index int)
	Set(index int,obj interface{})
	IndexOf(obj interface{}) int
}
