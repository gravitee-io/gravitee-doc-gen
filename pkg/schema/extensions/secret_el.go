package extensions

import (
	"encoding/json"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"strconv"
)

// TODO group that into gioConfig

const (
	SecretExtension = "x-secret"
	ElExtension     = "x-el"
)

type BoolValueCompiler struct {
	value bool
	Ext   string
}
type BoolValueSchema bool

func (c BoolValueCompiler) Compile(_ jsonschema.CompilerContext, m map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, ok := m[c.Ext]; ok {
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
