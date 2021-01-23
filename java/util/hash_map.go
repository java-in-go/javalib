package util

import (
	"reflect"
	"strconv"
)

type HashMap struct {
	m    map[interface{}]interface{}
	size int
}

func NewHashMap() *HashMap {
	return &HashMap{
		m:    make(map[interface{}]interface{}),
		size: 0,
	}
}
func NewHashMapCap(initialCapacity int) *HashMap {
	if initialCapacity < 0 {
		panic("Illegal initial capacity: " + strconv.Itoa(initialCapacity))
	}
	return &HashMap{
		m:    make(map[interface{}]interface{}, initialCapacity),
		size: initialCapacity,
	}
}
func (h *HashMap) Size() int {
	return h.size
}

func (h *HashMap) IsEmpty() bool {
	return h.size == 0
}

func (h *HashMap) ContainsKey(key interface{}) bool {
	_, exists := h.m[key]
	return exists
}

func (h *HashMap) ContainsValue(value interface{}) bool {
	for _, v := range h.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (h *HashMap) Get(key interface{}) interface{} {
	return h.m[key]
}

func (h *HashMap) Put(key interface{}, value interface{}) (interface{}, bool) {
	v, exists := h.m[key]
	h.m[key] = value
	h.size++
	return v, exists
}

func (h *HashMap) Remove(key interface{}) (interface{}, bool) {
	v, exists := h.m[key]
	if exists {
		delete(h.m, key)
		h.size--
	}
	return v, exists
}

func (h *HashMap) PutAll(m Map) {
	m.ForEach(func(k interface{}, v interface{}) {
		h.m[k] = v
		h.size++
	})
}

func (h *HashMap) Clear() {
	for k := range h.m {
		delete(h.m, k)
	}
	h.size = 0
}
func (h *HashMap) Keys() List {
	l := NewArrayListCap(len(h.m))
	for k := range h.m {
		l.Add(k)
	}
	return l
}
func (h *HashMap) Values() List {
	l := NewArrayListCap(len(h.m))
	for _, v := range h.m {
		l.Add(v)
	}
	return l
}
func (h *HashMap) ForEach(f func(interface{}, interface{})) {
	for k, v := range h.m {
		f(k, v)
	}
}
