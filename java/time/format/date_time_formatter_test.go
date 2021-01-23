package format

import (
	"reflect"
	"testing"
	"time"
)

func TestDateTimeFormatter_Format(t *testing.T) {
	type fields struct {
		pattern string
	}
	// 2020-11-22 11:56:03:00

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "yyyyMMddHHmmss",
			fields: fields{"yyyyMMddHHmmss"},
			want:   "20200102010603",
		},
		{
			name:   "yyyy-MM-dd-HH-mm-ss",
			fields: fields{"yyyy-MM-dd-HH-mm-ss"},
			want:   "2020-01-02-01-06-03",
		},
		{
			name:   "yyyy-MM-dd-HH-mm-ss",
			fields: fields{"yyyy-MM-dd-HH-mm-ss"},
			want:   "2020-01-02-01-06-03",
		},
		{
			name:   "yy-MM-dd-HH-mm-ss",
			fields: fields{"yy-MM-dd-HH-mm-ss"},
			want:   "20-01-02-01-06-03",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format := OfPattern(tt.fields.pattern)
			if got := format.Format(time.Date(2020, 01, 02, 01, 06, 03, 0, time.Local)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfPattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "yyyyMMdd", args: args{"yyyyMMdd"}, want: "20060102"},
		{name: "yMMdd", args: args{"yMMdd"}, want: "20060102"},
		{name: "yyMMdd", args: args{"yyMMdd"}, want: "060102"},
		{name: "yyyyMMMdd", args: args{"yyyyMMMdd"}, want: "2006Jan02"},
		{name: "yyyyMMMMdd", args: args{"yyyyMMMMdd"}, want: "2006January02"},
		{name: "yyyyMMd", args: args{"yyyyMMd"}, want: "2006012"},
		{name: "yyyyMMddHHmmss", args: args{"yyyyMMddHHmmss"}, want: "20060102150405"},
		{name: "yyyyMMddhhmmss", args: args{"yyyyMMddhhmmss"}, want: "20060102030405"},
		{name: "yyyyMMddhmmss", args: args{"yyyyMMddhmmss"}, want: "2006010230405"},
		{name: "yyyyMMddHHmmss", args: args{"yyyyMMddHHmmss"}, want: "20060102150405"},
		{name: "HHmmss", args: args{"HHmmss"}, want: "150405"},
		{name: "hhmss", args: args{"hhmss"}, want: "03405"},
		{name: "HHmms", args: args{"HHmms"}, want: "15045"},
		{name: "Hms", args: args{"Hms"}, want: "1545"},
		{name: "'yyyy'MMddHHmmss", args: args{"'yyyy'MMddHHmmss"}, want: "yyyy0102150405"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OfPattern(tt.args.pattern).GoPattern(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
