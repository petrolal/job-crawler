package classifier

import "strings"

func IsRemote(text string) bool {
	t := strings.ToUpper(text)

	keywords := []string{
		"REMOTO",
		"REMOTE",
		"HOME OFFICE",
		"TRABALHO REMOTO",
		"100% REMOTO",
		"ANYWHERE",
	}

	for _, k := range keywords {
		if strings.Contains(t, k) {
			return true
		}
	}
	return false
}

func IsBrazil(title, desc, location string) bool {
	t := strings.ToUpper(title + " " + desc + " " + location)

	locationKeywords := []string{
		"BRASIL",
		"BRAZIL",
		"BRASÍLIA",
		"DISTRITO FEDERAL",
		"DF",
	}

	hasLocation := false
	for _, l := range locationKeywords {
		if strings.Contains(t, l) {
			hasLocation = true
			break
		}
	}

	if !hasLocation {
		return false
	}

	return true
}
