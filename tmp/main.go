package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var b string

	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}

		lcs(a, b)
	}
}

func lcs(a, b string) {

	lenA, lenB := len(a), len(b)
	dp := make([][]int, lenA+1)

	for i := range dp {
		dp[i] = make([]int, lenB+1)
	}

	for i, ca := range a {
		for j, cb := range b {
			if ca == cb {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	fmt.Println(dp[lenA][lenB])

	i, j := lenA, lenB
	resBd := strings.Builder{}
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			resBd.WriteByte(a[i-1])
			i--
			j--
		} else if dp[i][j] <= dp[i-1][j] {
			i--
		} else {
			j--
		}
	}

	res := resBd.String()
	for index := len(res) - 1; index >= 0; index-- {
		fmt.Printf("%c", res[index])
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
