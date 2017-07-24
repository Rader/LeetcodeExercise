package interview

import "testing"

func Test_inversions(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal", args: args{vals: []int{2, 5, 3, 1, 10}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversions(tt.args.vals); got != tt.want {
				t.Errorf("inversions() = %v, want %v", got, tt.want)
			}
		})
	}
}
