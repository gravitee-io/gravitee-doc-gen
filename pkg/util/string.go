package util

import "strings"

func Title(value string) string {
	if value != "" {
		return strings.ToTitle(value[:1]) + value[1:]
	}
	return ""
}
