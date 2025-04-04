package util

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"text/template"
)

func RenderTemplateFromFile(file string, data interface{}) ([]byte, error) {
	tpl, err := CompileTemplateWithFunctions(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot parse template %s: %s", file, err.Error()))
	}
	return RenderTemplate(tpl, data)
}

func CompileTemplateWithFunctions(file string) (*template.Template, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return template.New(file).Funcs(template.FuncMap{
		"default": defaultTo,
		"asCheck": asCheck,
		"title":   Title,
		"code":    code}).Parse(string(content))
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
	if value == nil || reflect.ValueOf(value).IsZero() {
		return fallback
	}
	return value
}

func asCheck(value any) any {
	if isTrue, ok := value.(bool); ok {
		if isTrue {
			return "✅"
		} else {
			return " "
		}
	}
	return value
}

func code(value any) any {
	return fmt.Sprintf("`%v`", value)
}
