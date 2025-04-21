package chunks

import (
	"text/template"
)

type Consumable struct {
	Id     string
	Exists bool
}

type Processed struct {
	Data any
}

type Ready struct {
	Consumable
	Processed
	CompiledTemplate *template.Template
}

type Generated struct {
	Consumable
	Content string
}
