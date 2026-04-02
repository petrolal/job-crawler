package scoring

import "strings"

func Score(title, desc string, isRemote, isHybrid bool) int {
	text := strings.ToUpper(title + " " + desc)

	score := 1 // score base

	if containsAny(text,
		"QA",
		"QUALITY",
		"TEST",
		"TESTING",
		"ENGINEER IN TEST",
		"QUALITY ENGINEER",
		"TEST ENGINEER",
		"TEST ANALYST",
		"SDET",
	) {
		score += 3
	}

	if containsAny(text,
		"AUTOMATION",
		"SELENIUM",
		"CYPRESS",
		"PLAYWRIGHT",
	) {
		score += 2
	}

	if isRemote {
		score += 2
	}

	if isHybrid {
		score += 2
	}

	return score
}

func containsAny(text string, words ...string) bool {
	for _, w := range words {
		if strings.Contains(text, w) {
			return true
		}
	}
	return false
}
