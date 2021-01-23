package util

type Map interface {
	// Returns the number of key-value mappings in this map
	Size() int
	// Returns true if this map contains no key-value mappings.
	IsEmpty() bool
	// Returns true if this map contains a mapping for the specified key.
	ContainsKey(key interface{}) bool
	// Returns true  if this map maps one or more keys to the specified value.
	ContainsValue(value interface{}) bool
	// Returns the value to which the specified key is mapped,
	//or nil if this map contains no mapping for the key.
	Get(key interface{}) interface{}
	//  Associates the specified value with the specified key in this map
	//  (optional operation).  If the map previously contained a mapping for
	//  the key, the old value is replaced by the specified value
	// Returns old value,true  if this map contains a mapping for the specified key.
	Put(key interface{}, value interface{}) (interface{}, bool)
	// Removes the mapping for a key from this map if it is present
	// Returns old value,true  if this map contains a mapping for the specified key.
	Remove(key interface{}) (interface{}, bool)
	// Copies all of the mappings from the specified map to this map
	PutAll(m Map)
	// Removes all of the mappings from this map
	Clear()
	// Returns a List view of the keys contained in this map.
	Keys() List
	// Returns a List view of the values contained in this map.
	Values() List

	ForEach(f func(interface{}, interface{}))
}
