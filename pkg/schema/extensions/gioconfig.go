package extensions

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

const (
	GioConfigExtension = "gioConfig"
)

type GioConfigCompiler struct {
}
type GioConfigSchema struct {
	El     bool `json:"el"`
	Secret bool `json:"secret"`
}

func (c GioConfigCompiler) Compile(_ jsonschema.CompilerContext, m map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, ok := m[GioConfigExtension]; ok {
		var m map[string]interface{}
		if c, ok := e.(map[string]interface{}); !ok {
			return nil, errors.New(GioConfigExtension + " must be a map")
		} else {
			m = c
		}
		u := core.Unstructured(m)
		gioConfig, err := util.AnyMapToStruct[GioConfigSchema](&u)
		if err != nil {
			return nil, errors.New(GioConfigExtension + " cannot be parse json extension struct, check your gioConfig node")
		}
		return gioConfig, nil
	}
	return nil, nil
}

func (c GioConfigSchema) Validate(_ jsonschema.ValidationContext, _ interface{}) error {
	// we don't validate payloads, no implementation required
	return nil
}
