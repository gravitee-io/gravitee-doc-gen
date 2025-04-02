package chunks

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"text/template"
)

type Consumable struct {
	Id     string
	Exists bool
}

type TypeHandlerFunc func(configChunk config.Chunk) (Processed, error)
type TypeValidatorFunc func(configChunk config.Chunk) (bool, error)

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
