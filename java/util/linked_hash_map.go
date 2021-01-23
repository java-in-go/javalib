package util

import (
	"reflect"
)

const LinkedHashMapDefaultCapacity = 10

type LinkedHashMap struct {
	dataMap   map[interface{}]interface{}
	dataSlice []interface{}
	size      int
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		dataMap:   make(map[interface{}]interface{}),
		dataSlice: make([]interface{}, 0, LinkedHashMapDefaultCapacity),
		size:      0,
	}
}

func (m *LinkedHashMap) Size() int {
	return m.size
}

func (m *LinkedHashMap) IsEmpty() bool {
	return m.size == 0
}

func (m *LinkedHashMap) ContainsKey(key interface{}) bool {
	_, exists := m.dataMap[key]
	return exists
}

func (m *LinkedHashMap) ContainsValue(value interface{}) bool {
	for _, v := range m.dataMap {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (m *LinkedHashMap) Get(key interface{}) interface{} {
	return m.dataMap[key]
}

func (m *LinkedHashMap) Put(key interface{}, value interface{}) (interface{}, bool) {
	m.dataSlice = append(m.dataSlice, key)
	oldValue, exists := m.dataMap[key]
	m.dataMap[key] = value
	m.size++
	return oldValue, exists
}

func (m *LinkedHashMap) Remove(key interface{}) (interface{}, bool) {
	v, exists := m.dataMap[key]
	if exists {
		delete(m.dataMap, key)
		for i, v := range m.dataSlice {
			if reflect.DeepEqual(key, v) {
				m.dataSlice[i] = nil
			}
		}
		m.size--
	}
	return v, exists
}

func (m *LinkedHashMap) PutAll(maps Map) {
	maps.ForEach(func(k interface{}, v interface{}) {
		m.dataSlice = append(m.dataSlice, k)
		m.dataMap[k] = v
		m.size++
	})
}
func (m *LinkedHashMap) Clear() {
	l := m.dataSlice
	for i, v := range l {
		delete(m.dataMap, v)
		l[i] = nil
	}
	m.size = 0
}

func (m *LinkedHashMap) Keys() List {
	slice := m.dataSlice
	l := NewArrayListCap(m.size)
	for _, v := range slice {
		l.Add(v)
	}
	return l
}

func (m *LinkedHashMap) Values() List {
	l2 := NewArrayListCap(m.size)
	for _, v := range m.dataMap {
		l2.Add(v)
	}
	return l2
}

func (m *LinkedHashMap) ForEach(f func(interface{}, interface{})) {
	for k, v := range m.dataMap {
		f(k, v)
	}
}
