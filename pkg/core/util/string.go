// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"strings"
	"unicode"
)

func AnyToString(v any) string {
	return fmt.Sprintf("%v", v)
}

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
