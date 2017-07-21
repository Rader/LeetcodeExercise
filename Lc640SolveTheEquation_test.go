package LeetcodeExercise

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_parseItem(t *testing.T) {
	buf1 := bytes.Buffer{}
	buf1.WriteString("+1")
	buf2 := bytes.Buffer{}
	buf2.WriteString("-1")
	buf3 := bytes.Buffer{}
	buf3.WriteString("+1x")
	type args struct {
		buf bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want *item
	}{
		{name: "+1", args: args{buf: buf1}, want: &item{exponential: 0, coefficient: 1}},
		{name: "-1", args: args{buf: buf2}, want: &item{exponential: 0, coefficient: -1}},
		{name: "+1x", args: args{buf: buf3}, want: &item{exponential: 1, coefficient: +1}},
	}
	for _, tt := range tests {
		if got := parseItem(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. parseItem() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "x+5-3+x=6+x-2", args: args{equation: "x+5-3+x=6+x-2"}, want: "x=2"},
		{name: "x=x", args: args{equation: "x=x"}, want: "Infinite solutions"},
		{name: "2x=x", args: args{equation: "2x=x"}, want: "x=0"},
		{name: "2x+3x-6x=x+2", args: args{equation: "2x+3x-6x=x+2"}, want: "x=-1"},
		{name: "x=x+2", args: args{equation: "x=x+2"}, want: "No solution"},
		{name: "-x=-1", args: args{equation: "-x=-1"}, want: "x=1"},
	}
	for _, tt := range tests {
		if got := solveEquation(tt.args.equation); got != tt.want {
			t.Errorf("%q. solveEquation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
