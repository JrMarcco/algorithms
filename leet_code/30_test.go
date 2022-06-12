package leet_code

import "testing"

func findSubstring(s string, words []string) []int {
	wordsMap := map[string]int{}

	totalLen := 0
	for _, word := range words {
		totalLen += len(word)
		wordsMap[word]++
	}

	var ans []int
	for i := 0; i <= len(s)-totalLen; i++ {
		if valid(s[i:i+totalLen], words, wordsMap) {
			ans = append(ans, i)
		}
	}

	return ans
}

func valid(s string, words []string, wordsMap map[string]int) bool {
	k := len(s) / len(words)

	sMap := map[string]int{}
	for i := 0; i < len(s); i += k {
		sMap[s[i:i+k]]++
	}
	return equalsMap(sMap, wordsMap)
}

func equalsMap(a map[string]int, b map[string]int) bool {
	for key, valueA := range a {
		if valueB, ok := b[key]; ok && valueA == valueB {
			continue
		}
		return false
	}
	return true
}

func TestFindSubstring(t *testing.T) {
	t.Log(
		findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}),
	)
}
