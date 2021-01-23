package util

import "testing"

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

func TestStrFormat(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "123{}",
			args: args{
				format: "123{}",
				params: []interface{}{1, "dd", "23"},
			},
			want: "1231",
		},
		{
			name: "123{}567{sd}",
			args: args{
				format: "123{}5{}6{}7{sd}",
				params: []interface{}{1, "dd", "23", "dsfs"},
			},
			want: "12315dd6237{sd}",
		},
		{
			name: "123{sdf}",
			args: args{
				format: "123{sdf}",
				params: []interface{}{1, "dd", "23"},
			},
			want: "123{sdf}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrFormat(tt.args.format, tt.args.params...); got != tt.want {
				t.Errorf("StrFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringParam(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1}, want: "1"},
		{name: "2.1", args: args{2.1}, want: "2.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringParam(tt.args.v); got != tt.want {
				t.Errorf("getStringParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
