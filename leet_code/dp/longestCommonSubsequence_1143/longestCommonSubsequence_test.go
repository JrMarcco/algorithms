package longestCommonSubsequence_1143

func longestCommonSubsequence(text1 string, text2 string) int {
	lenA, lenB := len(text1), len(text2)
	dp := make([][]int, lenA+1)

	for i := range dp {
		dp[i] = make([]int, lenB+1)
	}

	for i, ca := range text1 {
		for j, cb := range text2 {
			if ca == cb {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[lenA][lenB]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
