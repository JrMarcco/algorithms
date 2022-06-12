package leet_code

import (
	"testing"
)

func groupAnagrams(strs []string) [][]string {

	strMap := map[[26]int][]string{}

	for _, str := range strs {
		letterCount := [26]int{}
		for _, ch := range str {
			letterCount[ch-'a']++
		}
		strMap[letterCount] = append(strMap[letterCount], str)
	}

	var ans [][]string
	for _, value := range strMap {
		ans = append(ans, value)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {

	t.Log(
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
	)
}
