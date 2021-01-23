package util

import (
	"reflect"
	"testing"
)

func TestLinkedHashMap_Clear(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	m := NewLinkedHashMap()
	m.Put("1", "2")
	m.Put("2", "3")
	m.Put("4", "2")
	m.Remove(2)
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			m.Clear()
			if got := m.Size(); got != 0 {
				t.Errorf("Clear() = %v, want %v", got, 0)
			}
		})
	}
}

func TestLinkedHashMap_ContainsKey(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		key interface{}
	}
	m := NewLinkedHashMap()
	m.Put("1", "2")
	m.Put("2", "3")
	m.Put("4", "2")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"1"},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"2"},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"3"},
			want:   false,
		},
		{
			name:   "4",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"4"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_ContainsValue(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		value interface{}
	}
	m := NewLinkedHashMap()
	m.Put("1", "2")
	m.Put("2", "3")
	m.Put("4", "5")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"1"},
			want:   false,
		},
		{
			name:   "2",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"2"},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"3"},
			want:   true,
		},
		{
			name:   "4",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"5"},
			want:   true,
		},
		{
			name:   "5",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{"6"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.ContainsValue(tt.args.value); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_ForEach(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		f func(interface{}, interface{})
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "4",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args: args{func(k interface{}, v interface{}) {
				if got := k.(int) + 1; got != v.(int) {
					t.Errorf("ForEach() = %v, want %v", got, v.(int))
				}
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			m.ForEach(tt.args.f)
		})
	}
}

func TestLinkedHashMap_Get(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		key interface{}
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{1},
			want:   2,
		},
		{
			name:   "2",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{2},
			want:   nil,
		},
		{
			name:   "3",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{3},
			want:   4,
		},
		{
			name:   "4",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{4},
			want:   nil,
		},
		{
			name:   "5",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{5},
			want:   6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_IsEmpty(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_Keys(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	l := NewArrayList()
	l.Add(1)
	l.Add(3)
	l.Add(5)
	tests := []struct {
		name   string
		fields fields
		want   List
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			want:   l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_Put(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	type want struct {
		exist bool
		value interface{}
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{1, 7},
			want:   want{true, 2},
		},
		{
			name:   "2",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{3, 3},
			want:   want{true, 3},
		},
		{
			name:   "3",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{4, 4},
			want:   want{false, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got, exists := m.Put(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want.value) && !reflect.DeepEqual(exists, tt.want.exist) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_PutAll(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		maps Map
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	keys := NewArrayList()
	keys.Add(1)
	keys.Add(3)
	keys.Add(5)
	values := NewArrayList()
	values.Add(2)
	values.Add(4)
	values.Add(6)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			args:   args{m},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewLinkedHashMap()
			m.PutAll(tt.args.maps)
			if !reflect.DeepEqual(m.Keys(), keys) && !reflect.DeepEqual(m.Values(), values) {
				t.Errorf("PutAll() = %v, want %v", m.Keys(), keys)
				t.Errorf("PutAll() = %v, want %v", m.Values(), values)
			}
		})
	}
}

func TestLinkedHashMap_Remove(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	type args struct {
		key interface{}
	}
	type want struct {
		value  interface{}
		exists bool
	}

	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "1",
			fields: fields{
				dataMap:   m.dataMap,
				dataSlice: m.dataSlice,
				size:      m.size,
			},
			args: args{1},
			want: want{2, true},
		},
		{
			name: "2",
			fields: fields{
				dataMap:   m.dataMap,
				dataSlice: m.dataSlice,
				size:      m.size,
			},
			args: args{2},
			want: want{nil, false},
		},
		{
			name: "3",
			fields: fields{
				dataMap:   m.dataMap,
				dataSlice: m.dataSlice,
				size:      m.size,
			},
			args: args{3},
			want: want{4, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got, exists := m.Remove(tt.args.key); !reflect.DeepEqual(got, tt.want.value) && !reflect.DeepEqual(exists, tt.want.exists) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_Size(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedHashMap_Values(t *testing.T) {
	type fields struct {
		dataMap   map[interface{}]interface{}
		dataSlice []interface{}
		size      int
	}
	m := NewLinkedHashMap()
	m.Put(1, 2)
	m.Put(3, 4)
	m.Put(5, 6)
	value := NewArrayList()
	value.Add(2)
	value.Add(4)
	value.Add(6)
	tests := []struct {
		name   string
		fields fields
		want   List
	}{
		{
			name:   "1",
			fields: fields{m.dataMap, m.dataSlice, m.size},
			want:   value,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LinkedHashMap{
				dataMap:   tt.fields.dataMap,
				dataSlice: tt.fields.dataSlice,
				size:      tt.fields.size,
			}
			if got := m.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkedHashMap(t *testing.T) {
	mappp := NewLinkedHashMap()
	tests := []struct {
		name string
		want *LinkedHashMap
	}{
		{name: "1", want: mappp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedHashMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
