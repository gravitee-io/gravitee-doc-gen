package schema

import (
	"encoding/json"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"strconv"
)

const (
	SecretExtension = "x-secret"
	elExtension     = "x-el"
)

type boolValueCompiler struct {
	value bool
	ext   string
}
type BoolValueSchema bool

func (c boolValueCompiler) Compile(_ jsonschema.CompilerContext, m map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, ok := m[c.ext]; ok {
		switch v := e.(type) {
		case bool:
			return BoolValueSchema(v), nil
		case string:
			boolVal, err := strconv.ParseBool(v)
			return BoolValueSchema(boolVal), err
		case json.Number:
			i, err := v.Int64()
			return BoolValueSchema(i > 0), err
		}
	}
	return nil, nil
}

func (c BoolValueSchema) Validate(_ jsonschema.ValidationContext, _ interface{}) error {
	// we don't validate payloads, no implementation required
	return nil
}

func Compile(schemaFile string) (*jsonschema.Schema, error) {

	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft7
	compiler.ExtractAnnotations = true
	compiler.RegisterExtension(SecretExtension, nil, &boolValueCompiler{ext: SecretExtension})
	compiler.RegisterExtension(elExtension, nil, &boolValueCompiler{ext: elExtension})

	return compiler.Compile(schemaFile)

}
