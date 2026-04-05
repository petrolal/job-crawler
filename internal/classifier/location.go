// Package classifier provides functions to classify job postings
// based on location and work arrangement signals.
package classifier

import (
	"regexp"
	"strings"
)

var dfWord = regexp.MustCompile(`\bDF\b`)

// IsRemote reports whether the given text contains signals indicating
// a remote work arrangement.
func IsRemote(text string) bool {
	t := strings.ToUpper(text)
	for _, k := range []string{"REMOTO", "REMOTE", "HOME OFFICE", "TRABALHO REMOTO", "100% REMOTO", "ANYWHERE"} {
		if strings.Contains(t, k) {
			return true
		}
	}
	return false
}

// IsBrazil reports whether the job is located in Brazil based on
// the title, description, and location fields.
// Broad keywords are matched against title and location only to avoid
// false positives in long descriptions. "DF" is matched as a whole word.
func IsBrazil(title, desc, location string) bool {
	loc := strings.ToUpper(location)
	full := strings.ToUpper(title + " " + location)
	for _, k := range []string{
		"BRASIL", "BRAZIL", "BRASÍLIA", "BRASILIA", "DISTRITO FEDERAL",
		"SÃO PAULO", "SAO PAULO", "RIO DE JANEIRO", "BELO HORIZONTE",
	} {
		if strings.Contains(full, k) {
			return true
		}
	}
	return dfWord.MatchString(loc)
}
