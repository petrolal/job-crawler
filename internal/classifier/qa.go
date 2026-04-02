package classifier

import "strings"

// IsLikelyQA identifica vagas QA mesmo sem "QA" explícito
func IsLikelyQA(title, desc string) bool {
	text := strings.ToUpper(title + " " + desc)

	signals := []string{
		"QA",
		"QUALITY ASSURANCE",
		"QUALITY ENGINEER",
		"TEST",
		"TESTING",
		"TEST ENGINEER",
		"ENGINEER IN TEST",
		"TEST ANALYST",
		"SDET",
		"AUTOMATION TEST",
		"TEST AUTOMATION",
	}

	for _, s := range signals {
		if strings.Contains(text, s) {
			return true
		}
	}
	return false
}
