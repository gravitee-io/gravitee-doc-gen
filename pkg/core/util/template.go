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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
)

func RenderTemplateFromFile(file string, data interface{}) ([]byte, error) {
	tpl, err := TemplateWithFunctions(file)
	if err != nil {
		return nil, fmt.Errorf("cannot parse template %s: %w", file, err)
	}
	return RenderTemplate(tpl, data)
}

func TemplateWithFunctions(file string) (*template.Template, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return template.New(file).Funcs(template.FuncMap{
		"default":    defaultTo,
		"ternary":    ternary,
		"indent":     indent,
		"pad":        pad,
		"quote":      quote,
		"icz":        increase,
		"joinset":    joinset,
		"mvmdheader": MoveMarkdownHeader,
		"title":      Title,
		"upper":      upper,
	},
	).Parse(string(content))
}

func RenderTemplate(tpl *template.Template, data interface{}) ([]byte, error) {
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	err := tpl.Execute(buffer, data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func defaultTo(value any, fallback any) any {
	if b, ok := value.(bool); ok {
		return b
	}
	if value == nil || reflect.ValueOf(value).IsZero() {
		return fallback
	}
	return value
}

func ternary(isTrue bool, ifTrue any, ifFalse any) any {
	if isTrue {
		return ifTrue
	} else {
		return ifFalse
	}
}

func indent(amount int, value any) string {
	scanner := bufio.NewScanner(strings.NewReader(AnyToString(value)))
	scanner.Split(bufio.ScanLines)

	buffer := bytes.Buffer{}

	// re-enter first line
	scanner.Scan()
	line := scanner.Text()
	buffer.WriteString(line)
	buffer.WriteString("\n")

	padding := strings.Repeat(" ", amount)
	for scanner.Scan() {
		buffer.WriteString(fmt.Sprintf("%s%s\n", padding, scanner.Text()))
	}

	b := buffer.Bytes()
	return string(b[:len(b)-1])
}

func quote(value any) any {
	if s, ok := value.(string); ok {
		return fmt.Sprintf(`"%s"`, s)
	}
	return value
}

func increase(value int) int {
	return value + 1
}

func pad(amount int) string {
	return strings.Repeat(" ", amount)
}

func joinset(set map[any]bool, separator string, surrounding string) string {
	items := make([]string, 0)
	for value := range set {
		if s, ok := value.(string); ok {
			items = append(items, fmt.Sprintf("%s%s%s", surrounding, s, surrounding))
		} else {
			items = append(items, AnyToString(value))
		}
	}
	return strings.Join(items, separator)
}

func MoveMarkdownHeader(offset int, content string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Split(bufio.ScanLines)
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			var found bool
			var actual = 0
			for line[0] == '#' {
				found = true
				line = line[1:]
				actual++
			}
			newLevel := strings.Repeat("#", actual+offset)
			if found {
				buffer.WriteString(newLevel)
			}
		}
		buffer.WriteString(line + "\n")
	}
	return buffer.String()
}

func upper(str any) string {
	return strings.ToUpper(AnyToString(str))
}
