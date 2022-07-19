package subdomainVisits_811

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var visitMap map[string]int

func subdomainVisits(cpdomains []string) []string {

	if cpdomains == nil || len(cpdomains) == 0 {
		return []string{}
	}

	visitMap = make(map[string]int)
	for _, cpdomain := range cpdomains {
		dealVisits(cpdomain)
	}

	if visitMap != nil {
		var ans []string
		for key, value := range visitMap {
			ans = append(ans, fmt.Sprintf("%d %s", value, key))
		}
		return ans
	}

	return []string{}
}

func dealVisits(cpdomain string) {
	strs := strings.Split(cpdomain, " ")
	visitTimes, _ := strconv.Atoi(strs[0])

	calcVisitTimes(strs[1], visitTimes)
}

func calcVisitTimes(domain string, visitTimes int) {
	index := strings.Index(domain, ".")

	if times, ok := visitMap[domain]; ok {
		visitMap[domain] = times + visitTimes
	} else {
		visitMap[domain] = visitTimes
	}

	if index == -1 {
		return
	}

	subDomain := domain[index+1:]
	calcVisitTimes(subDomain, visitTimes)
}

func TestSubdomainVisits(t *testing.T) {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	t.Log(subdomainVisits(cpdomains))
}
