package util

import (
	"strings"
	"unicode"
)

func Title(value any) string {
	str := AnyToString(value)
	if value != "" {
		return strings.ToTitle(str[:1]) + str[1:]
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
