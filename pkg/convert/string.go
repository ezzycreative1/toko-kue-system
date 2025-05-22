package convert

import (
	"strings"
	"unicode"
)

func ReplaceSpaceAndDot(str string, replacement rune) string {
	modifiedStr := strings.Map(func(r rune) rune {
		if r == ' ' || r == '.' {
			return replacement
		}
		return unicode.ToLower(r) // Mengubah karakter menjadi huruf kecil
	}, str)
	return modifiedStr
}
