package util

import (
	"strings"
	"unicode"
)

func Title(value string) string {
	if value != "" {
		return strings.ToTitle(value[:1]) + value[1:]
	}
	return ""
}

func TitleCaseToTitle(value string) string {
	runes := []rune(value)
	output := make([]rune, 0, len(runes))
	output = append(output, runes[0])
	for i := 1; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsUpper(r) {
			output = append(output, ' ')
		}
		output = append(output, r)
	}
	return string(output)
}
