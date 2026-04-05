package classifier

import (
	"regexp"
	"strings"
)

var dfWord = regexp.MustCompile(`\bDF\b`)

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
	loc := strings.ToUpper(location)
	full := strings.ToUpper(title + " " + location)

	broadKeywords := []string{
		"BRASIL", "BRAZIL", "BRASÍLIA", "BRASILIA", "DISTRITO FEDERAL",
		"SÃO PAULO", "SAO PAULO", "RIO DE JANEIRO", "BELO HORIZONTE",
	}
	for _, k := range broadKeywords {
		if strings.Contains(full, k) {
			return true
		}
	}

	return dfWord.MatchString(loc)
}
