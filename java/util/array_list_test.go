package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		obj interface{}
	}
	type Obj struct {
		data string
	}
	obj1 := &Obj{"1"}
	obj2 := &Obj{"2"}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{0, nil},
			args:   args{obj1},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{0, nil},
			args:   args{obj2},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Add(tt.args.obj); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Clear(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "1", fields: fields{}},
		{name: "2", fields: fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArrayList()
			a.Add(tt)
			if a.size != 1 {
				t.Errorf("Clear() = %v, want %v", a.size, 1)
			}
			a.Clear()
			if a.size != 0 {
				t.Errorf("Clear() = %v, want %v", a.size, 0)
			}
		})
	}
}

func TestArrayList_Contains(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		obj interface{}
	}
	list1 := NewArrayList()
	list1.Add(args{"1"})
	list1.Add(args{"2"})
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "EmptyList1", fields: fields{0, nil}, args: args{nil}, want: false},
		{name: "EmptyList2", fields: fields{0, nil}, args: args{nil}, want: false},
		{name: "NotEmptyList2", fields: fields{list1.size, list1.elementData}, args: args{"2"}, want: false},
		{name: "NotEmptyListArgs2", fields: fields{list1.size, list1.elementData}, args: args{args{"2"}}, want: true},
		{name: "NotEmptyList3", fields: fields{list1.size, list1.elementData}, args: args{"3"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Contains(tt.args.obj); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
	}
	list1 := NewArrayList()
	list1.Add(args{1})
	list1.Add(args{2})

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		/*{
			name:   "1",
			fields: fields{0,nil},
			args:   args{1},
			want:   nil,
		},*/
		/*{
			name:   "10",
			fields: fields{0,nil},
			args:   args{1},
			want:   nil,
		},*/
		{
			name:   "1",
			fields: fields{list1.size, list1.elementData},
			args:   args{1},
			want:   args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		obj interface{}
	}
	list1 := NewArrayList()
	list1.Add(args{"1"})
	list1.Add(args{"2"})
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1",
			fields: fields{0, nil},
			args:   args{1},
			want:   -1,
		}, {
			name:   "2",
			fields: fields{list1.size, list1.elementData},
			args:   args{args{"1"}},
			want:   0,
		}, {
			name:   "3",
			fields: fields{list1.size, list1.elementData},
			args:   args{args{"2"}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.IndexOf(tt.args.obj); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		obj interface{}
	}
	list1 := NewArrayList()
	list1.Add(args{"1"})
	list1.Add(args{"2"})
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty",
			fields: fields{0, nil},
			want:   true,
		},
		{
			name:   "notempty",
			fields: fields{list1.size, list1.elementData},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Remove(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
	}
	list1 := NewArrayList()
	list1.Add(args{1})
	list1.Add(args{2})
	t.Logf("size:%d,len:%d", list1.size, len(list1.elementData))
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		/*{
			name:   "1",
			fields: fields{0, nil},
			args:   args{1},
			want:   nil,
		},*/{
			name:   "2",
			fields: fields{list1.size, list1.elementData},
			args:   args{1},
			want:   args{2},
		}, {
			name:   "3",
			fields: fields{list1.size, list1.elementData},
			args:   args{0},
			want:   args{1},
		}, {
			name:   "4",
			fields: fields{list1.size, list1.elementData},
			args:   args{0},
			want:   args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
			t.Logf("size:%d,len:%d", a.size, len(a.elementData))
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
		obj   interface{}
	}
	list1 := NewArrayList()
	list1.Add(args{1, nil})
	list1.Add(args{2, nil})
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "1",
			fields: fields{list1.size, list1.elementData},
			args:   args{1, args{1, nil}},
			want:   args{2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Set(tt.args.index, tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Size(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
	}
	list1 := NewArrayList()
	list1.Add(args{1})
	list1.Add(args{2})
	tests := []struct {
		name   string
		fields fields
		want   int
	}{

		{
			name:   "1",
			fields: fields{0, nil},
			want:   0,
		},
		{
			name:   "2",
			fields: fields{list1.Size(), list1.elementData},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_ensureCapacityInternal(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{0, EmptyElementData},
			args:   args{1},
			want:   10,
		},
		{
			name:   "2",
			fields: fields{0, EmptyElementData},
			args:   args{10},
			want:   10,
		},
		{
			name:   "3",
			fields: fields{0, EmptyElementData},
			args:   args{11},
			want:   11,
		},
		{
			name:   "4",
			fields: fields{0, EmptyElementData},
			args:   args{-11},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			if got := a.ensureCapacityInternal(tt.args.index); got != tt.want {
				t.Errorf("ensureCapacityInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_ensureExplicitCapacity(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		minCapacity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{0, nil},
			args:   args{1},
			want:   1,
		},
		{
			name:   "2",
			fields: fields{10, nil},
			args:   args{2},
			want:   2,
		},
		{
			name:   "3",
			fields: fields{1, nil},
			args:   args{1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			a.ensureExplicitCapacity(tt.args.minCapacity)
			got := cap(a.elementData)
			if got != tt.want {
				t.Errorf("ensureExplicitCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_rangeCheck(t *testing.T) {
	type fields struct {
		size        int
		elementData []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				size:        tt.fields.size,
				elementData: tt.fields.elementData,
			}
			fmt.Println(a)
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1,2", args: args{1, 2}, want: 2},
		{name: "5,5", args: args{5, 5}, want: 5},
		{name: "5,2", args: args{5, 2}, want: 5},
		{name: "-1,0", args: args{-1, 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArrayList(t *testing.T) {
	tests := []struct {
		name string
		want *ArrayList
	}{
		{name: "1", want: NewArrayList()},
		{name: "2", want: NewArrayList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArrayListCap(t *testing.T) {
	type args struct {
		initialCapacity int
	}
	tests := []struct {
		name string
		args args
		want *ArrayList
	}{
		{name: "1", args: args{
			1,
		}, want: NewArrayListCap(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayListCap(tt.args.initialCapacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayListCap() = %v, want %v", got, tt.want)
			}
		})
	}
}
