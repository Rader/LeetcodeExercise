package interview

import (
	"reflect"
	"testing"
)

func Test_findMin(t *testing.T) {
	root := BuildBST([]int{8, 9, 4, 3, 6, 5, 7})
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "min is left leaf", args: args{root: root}, want: 3},
		{name: "min in right child", args: args{root: root.L.R}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.root); got.V != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildBST(t *testing.T) {
	bst := &TreeNode{V: 2}
	bst.L = &TreeNode{V: 1}
	bst.R = &TreeNode{V: 3}
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "normal", args: args{vals: []int{2, 1, 3}}, want: bst},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildBST(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_delBSTNodeDoubleLinked(t *testing.T) {
	{
		want := BuildBSTDoubleLinked([]int{8, 9, 4, 6, 5, 7})
		t.Run("delete left leaf", func(t *testing.T) {
			root := BuildBSTDoubleLinked([]int{8, 9, 4, 3, 6, 5, 7})
			n := findNodeDoubleLinked(root, 3)
			if got := delBSTNodeDoubleLinked(root, n); !reflect.DeepEqual(got, want) {
				t.Errorf("delBSTNodeDoubleLinked() = %v, want %v", got, want)
			}
		})
	}
	{
		want := BuildBSTDoubleLinked([]int{8, 9, 5, 3, 6, 7})
		t.Run("delete middle has right child", func(t *testing.T) {
			root := BuildBSTDoubleLinked([]int{8, 9, 4, 3, 6, 5, 7})
			n := findNodeDoubleLinked(root, 4)
			if got := delBSTNodeDoubleLinked(root, n); !reflect.DeepEqual(got, want) {
				t.Errorf("delBSTNodeDoubleLinked() = %v, want %v", got, want)
			}
		})
	}
	{
		want := BuildBSTDoubleLinked([]int{9, 5, 3, 6, 7})
		t.Run("delete root", func(t *testing.T) {
			root := BuildBSTDoubleLinked([]int{8, 9, 4, 3, 6, 5, 7})
			n := findNodeDoubleLinked(root, 8)
			if got := delBSTNodeDoubleLinked(root, n); !reflect.DeepEqual(got, want) {
				t.Errorf("delBSTNodeDoubleLinked() = %v, want %v", got, want)
			}
		})
	}
}

func Test_findNode(t *testing.T) {
	root := BuildBST([]int{8, 9, 4, 3, 6, 5, 7})
	type args struct {
		root *TreeNode
		v    int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "exist", args: args{root: root, v: 3}, want: &TreeNode{V: 3}},
		{name: "not exist", args: args{root: root, v: 1}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNode(tt.args.root, tt.args.v); got != tt.want && got.V != tt.want.V {
				t.Errorf("findNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
