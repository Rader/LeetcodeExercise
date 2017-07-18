package LeetcodeExercise

/** URL: https://leetcode.com/problems/delete-operation-for-two-strings
Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.

Example 1:
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Note:
The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.
*/

func minDistance(word1 string, word2 string) int {
	lsub := lenOfLongestCommonSubsequence(word1, word2)
	//total steps = steps change word1 to common subsequence + steps change word2 to common subsequence
	return (len(word1) - lsub) + (len(word2) - lsub)
}

//longest common subsequence does not need to be continuous
func lenOfLongestCommonSubsequence(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	if l1 == 0 || l2 == 0 {
		return 0
	}

	if s1[l1-1] == s2[l2-1] {
		return lenOfLongestCommonSubsequence(s1[:l1-1], s2[:l2-1]) + 1
	}
	return max(lenOfLongestCommonSubsequence(s1[:l1-1], s2), lenOfLongestCommonSubsequence(s1, s2[:l2-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
