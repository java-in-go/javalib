package util

type List interface {
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

	ForEach(f func(value interface{}))
	ForEachIndex(f func(index int, value interface{}))
	Map(f func(value interface{}) interface{}) List
	Reduce(f func(v1 interface{}, v2 interface{}) interface{}) interface{}
}
