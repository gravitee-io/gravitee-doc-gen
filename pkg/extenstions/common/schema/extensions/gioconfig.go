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

package extensions

import (
	"errors"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
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

func (c GioConfigCompiler) Compile(
	_ jsonschema.CompilerContext,
	schema map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, exists := schema[GioConfigExtension]; exists {
		var config map[string]interface{}
		if c, ok := e.(map[string]interface{}); !ok {
			return nil, errors.New(GioConfigExtension + " must be a map")
		} else {
			config = c
		}
		u := util.Unstructured(config)
		gioConfig, err := util.AnyMapToStruct[GioConfigSchema](&u)
		if err != nil {
			return nil, errors.New(GioConfigExtension + " cannot be parse json extension struct, check your gioConfig node")
		}
		return gioConfig, nil
	}
	return GioConfigSchema{}, nil
}

func (c GioConfigSchema) Validate(_ jsonschema.ValidationContext, _ interface{}) error {
	// we don't validate payloads, no implementation required
	return nil
}
