package interview

//LongestCommonSubsequence get length of the longest common subsequence
//time complexity is O(n^2)
func LongestCommonSubsequence(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	if l1 == 0 || l2 == 0 {
		return 0
	}

	if s1[l1-1] == s2[l2-1] {
		return LongestCommonSubsequence(s1[:l1-1], s2[:l2-1]) + 1
	}
	return max(LongestCommonSubsequence(s1[:l1-1], s2), LongestCommonSubsequence(s1, s2[:l2-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
