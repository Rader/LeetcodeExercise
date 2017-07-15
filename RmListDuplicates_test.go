package interview

import (
	"reflect"
	"testing"
)

func TestRmListDuplicates(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "empty list", args: args{n: createList([]int{})}, want: nil},
		{name: "duplicates at beginning", args: args{n: createList([]int{2, 2, 2, 3, 4})}, want: createList([]int{3, 4})},
		{name: "duplicates at the end", args: args{n: createList([]int{0, 1, 2, 2, 2})}, want: createList([]int{0, 1})},
		{name: "all duplicates ", args: args{n: createList([]int{2, 2, 2})}, want: nil},
		{name: "duplicates in middle", args: args{n: createList([]int{3, 2, 2, 1})}, want: createList([]int{3, 1})},
	}
	for _, tt := range tests {
		if got := RmListDuplicates(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			printList(t, got)
			t.Errorf("%q. RmListDuplicates() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func createList(s []int) *Node {
	h := new(Node)
	tail := h
	for _, i := range s {
		n := &Node{V: i}
		tail.Next = n
		tail = tail.Next
	}
	return h.Next
}

func printList(t *testing.T, n *Node) {
	for ; n != nil; n = n.Next {
		t.Log(n.V)
	}
}
