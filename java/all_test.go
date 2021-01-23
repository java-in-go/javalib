package java

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  interface{}
		falseValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "1<2",
			args: args{
				1 < 2,
				3,
				4,
			},
			want: 3,
		},
		{
			name: "1+1=2",
			args: args{
				1+1 == 2,
				5,
				7,
			},
			want: 5,
		},
		{
			name: "1+1=3",
			args: args{
				1+1 == 3,
				2,
				3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewArrayList(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		{name: "New"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewArrayList()
			}
		})
	}
}
func BenchmarkNewArray(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		{name: "NewArray"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				i2 := make([]interface{}, 1, 10)
				i2[0] = 1
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
		{name: "1", args: args{1}, want: NewArrayListCap(1)},
		{name: "11", args: args{1}, want: NewArrayListCap(11)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayListCap(tt.args.initialCapacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayListCap() = %v, want %v", got, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMapCap(tt.args.initialCapacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMapCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkedHashMap(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedHashMap
	}{
		{name: "1", want: NewLinkedHashMap()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedHashMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedHashMap() = %v, want %v", got, tt.want)
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
