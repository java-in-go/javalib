package util

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		obj interface{}
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     bool
		wantSize int
	}{
		{
			name:     "1",
			fields:   fields{list.New(), 0},
			args:     args{1},
			want:     true,
			wantSize: 1,
		},
		{
			name:     "2",
			fields:   fields{list1.list, list1.size},
			args:     args{1},
			want:     true,
			wantSize: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Add(tt.args.obj); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
			size := a.Size()
			if size != tt.wantSize {
				t.Errorf("Add() = %v, want size %v", size, tt.wantSize)
			}
		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name     string
		fields   fields
		wantSize int
	}{
		{name: "1", fields: fields{list.New(), 0}, wantSize: 0},
		{name: "2", fields: fields{list1.list, list1.size}, wantSize: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			a.Clear()
			if a.Size() != tt.wantSize {
				t.Errorf("Clear() = %v, want size %v", a.Size(), tt.wantSize)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		obj interface{}
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "1", fields: fields{list1.list, list1.size}, args: args{1}, want: true},
		{name: "1", fields: fields{list1.list, list1.size}, args: args{2}, want: true},
		{name: "1", fields: fields{list1.list, list1.size}, args: args{3}, want: true},
		{name: "1", fields: fields{list1.list, list1.size}, args: args{4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Contains(tt.args.obj); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		index int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   2,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   3,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		obj interface{}
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   0,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   1,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.IndexOf(tt.args.obj); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			want:   false,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			want:   false,
		},
		{
			name:   "3",
			fields: fields{list.New(), 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		index int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   2,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   3,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveObj(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		obj interface{}
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.RemoveObj(tt.args.obj); got != tt.want {
				t.Errorf("RemoveObj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Set(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		index int
		obj   interface{}
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1, 1},
			want:   2,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2, 2},
			want:   3,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Set(tt.args.index, tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			want:   3,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			want:   3,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_checkElementIndex(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		index int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			a.checkElementIndex(tt.args.index)
		})
	}
}

func TestLinkedList_isElementIndex(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   true,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.isElementIndex(tt.args.index); got != tt.want {
				t.Errorf("isElementIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_node(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		index int
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args:   args{1},
			want:   2,
		},
		{
			name:   "2",
			fields: fields{list1.list, list1.size},
			args:   args{2},
			want:   3,
		},
		{
			name:   "3",
			fields: fields{list1.list, list1.size},
			args:   args{0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.node(tt.args.index); !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("node() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedList
	}{
		{name: "1", want: NewLinkedList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_ForEach(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		f func(interface{})
	}
	list1 := NewLinkedList()
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "1",
			fields: fields{list1.list, list1.size},
			args: args{f: func(value interface{}) {

			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			a.ForEach(func(v interface{}) {
				fmt.Println(v)
			})
		})
	}
}

func TestLinkedList_Map(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		f func(interface{}) interface{}
	}
	l := NewLinkedList()
	l.Add("1")
	l.Add("2")
	l.Add("3")
	l2 := NewLinkedList()
	l2.Add("1x")
	l2.Add("2x")
	l2.Add("3x")
	tests := []struct {
		name   string
		fields fields
		args   args
		want   List
	}{
		{
			name:   "1",
			fields: fields{l.list, l.size},
			args: args{f: func(value interface{}) interface{} {
				return value.(string) + "x"
			}},
			want: l2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Reduce(t *testing.T) {
	type fields struct {
		list *list.List
		size int
	}
	type args struct {
		f func(interface{}, interface{}) interface{}
	}
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{l.list, l.size},
			args: args{f: func(v1 interface{}, v2 interface{}) interface{} {
				return v1.(int) + v2.(int)
			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LinkedList{
				list: tt.fields.list,
				size: tt.fields.size,
			}
			if got := a.Reduce(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
