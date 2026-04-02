package extractor

import "regexp"

var (
	emailRegex = regexp.MustCompile(
		`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`,
	)

	emailWithContextRegex = regexp.MustCompile(
		`(.{0,40}[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}.{0,40})`,
	)
)

func ExtractEmailWithContext(text string) (email, context string) {
	match := emailWithContextRegex.FindString(text)
	if match == "" {
		return "", ""
	}
	return emailRegex.FindString(match), match
}

func ClassifyEmail(email string) string {
	switch {
	case regexp.MustCompile(`recrut|talento|career|job|people`).MatchString(email):
		return "recruitment"
	case regexp.MustCompile(`info|contact`).MatchString(email):
		return "generic"
	default:
		return "unknown"
	}
}
