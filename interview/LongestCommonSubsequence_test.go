package interview

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "normal", args: args{s1: "sea", s2: "eat"}, want: 2},
		{name: "normal2", args: args{s1: "seaeatp", s2: "earp"}, want: 3}, //common:eap
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
