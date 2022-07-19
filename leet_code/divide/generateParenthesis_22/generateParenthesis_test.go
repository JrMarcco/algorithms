package generateParenthesis_22

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var ans []string
	for k := 1; k <= n; k++ {
		a := generateParenthesis(k - 1)
		b := generateParenthesis(n - k)
		for _, strA := range a {
			for _, strB := range b {
				ans = append(ans, fmt.Sprintf("(%s)%s", strA, strB))
			}
		}
	}

	return ans
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}
