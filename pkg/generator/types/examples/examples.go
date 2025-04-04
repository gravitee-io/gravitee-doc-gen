package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
)

type Examples struct {
	current  int
	Snippets []Snippet
}

type Snippet struct {
	Title       string
	Description string
	Filename    string
	Language    Language
	Code        string
}

type Code struct {
	config.Plugin
	Properties    map[string]interface{}
	Configuration string
}
