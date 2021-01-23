package util

import (
	"reflect"
	"testing"
)

func TestHashMap_Clear(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "1", fields: fields{
			m:    mmm,
			size: 2,
		}},
		{name: "2", fields: fields{
			m:    nil,
			size: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			h.Clear()
			if got := h.Size(); got != 0 {
				t.Errorf("Clear() = %v, want %v", got, 0)
			}
		})
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	type args struct {
		key interface{}
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{"1"},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{mmm, 2},
			args:   args{"2"},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{mmm, 2},
			args:   args{"3"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_ContainsValue(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}

	type args struct {
		value interface{}
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{"1"},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{mmm, 2},
			args:   args{"2"},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{mmm, 2},
			args:   args{"3"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.ContainsValue(tt.args.value); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_ForEach(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	type args struct {
		f func(interface{}, interface{})
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "1",
			fields: fields{m: mmm, size: 2},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			h.ForEach(func(k interface{}, v interface{}) {
				if tt.fields.m[k] != v {
					t.Errorf("ForEach() = %v, want %v", tt.fields.m[k], v)
				}
			})
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	type args struct {
		key interface{}
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{"1"},
			want:   "1",
		},
		{
			name:   "2",
			fields: fields{mmm, 2},
			args:   args{"2"},
			want:   "2",
		},
		{
			name:   "3",
			fields: fields{mmm, 2},
			args:   args{"3"},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_IsEmpty(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			want:   false,
		},
		{
			name:   "1",
			fields: fields{nil, 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Keys(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	l := NewArrayList()
	l.Add("1")
	l.Add("2")
	tests := []struct {
		name   string
		fields fields
		want   List
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			want:   l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Put(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	type args struct {
		key   interface{}
		value interface{}
	}
	type want struct {
		value  interface{}
		exists bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{"1", "3"},
			want:   want{"1", true},
		},
		{
			name:   "2",
			fields: fields{mmm, 2},
			args:   args{"2", "2"},
			want:   want{"2", true},
		},
		{
			name:   "3",
			fields: fields{mmm, 2},
			args:   args{"3", "5"},
			want:   want{nil, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got, exists := h.Put(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want.value) && !reflect.DeepEqual(exists, tt.want.exists) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_PutAll(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	type args struct {
		m Map
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	mmm2 := make(map[interface{}]interface{})
	mmm2["1"] = "1"
	mmm2["2"] = "2"
	mmm2["4"] = "5"
	mmm2["2"] = "7"
	m := NewHashMap()
	m.Put("4", "5")
	m.Put("2", "7")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[interface{}]interface{}
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{m},
			want:   mmm2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			h.PutAll(tt.args.m)
			h.ForEach(func(k interface{}, v interface{}) {
				if got := mmm2[k]; !reflect.DeepEqual(got, v) {
					t.Errorf("PutAll() = %v, want %v", v, got)
				}
			})
		})
	}
}

func TestHashMap_Remove(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	type args struct {
		key interface{}
	}
	type want struct {
		value  interface{}
		exists bool
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   args{"1"},
			want:   want{"1", true},
		},
		{
			name:   "2",
			fields: fields{mmm, 2},
			args:   args{"2"},
			want:   want{"2", true},
		},
		{
			name:   "3",
			fields: fields{mmm, 2},
			args:   args{"3"},
			want:   want{nil, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got, exists := h.Remove(tt.args.key); !reflect.DeepEqual(got, tt.want.value) && !reflect.DeepEqual(exists, tt.want.exists) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Size(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "1"
	mmm["2"] = "2"
	tests := []struct {
		name   string
		fields fields
		args   map[interface{}]interface{}
		want   int
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			args:   mmm,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashMap()
			for k, v := range tt.args {
				h.Put(k, v)
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Values(t *testing.T) {
	type fields struct {
		m    map[interface{}]interface{}
		size int
	}
	mmm := make(map[interface{}]interface{})
	mmm["1"] = "3"
	mmm["2"] = "4"
	l := NewArrayList()
	l.Add("3")
	l.Add("4")
	tests := []struct {
		name   string
		fields fields
		want   List
	}{
		{
			name:   "1",
			fields: fields{mmm, 2},
			want:   l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				m:    tt.fields.m,
				size: tt.fields.size,
			}
			if got := h.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashMap(t *testing.T) {
	tests := []struct {
		name string
		want *HashMap
	}{
		{name: "1", want: NewHashMap()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashMapCap(t *testing.T) {
	type args struct {
		initialCapacity int
	}
	tests := []struct {
		name string
		args args
		want *HashMap
	}{
		{name: "1", args: args{1}, want: NewHashMapCap(1)},
		{name: "1", args: args{2}, want: NewHashMapCap(2)},
		{name: "1", args: args{3}, want: NewHashMapCap(3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMapCap(tt.args.initialCapacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMapCap() = %v, want %v", got, tt.want)
			}
		})
	}
}
