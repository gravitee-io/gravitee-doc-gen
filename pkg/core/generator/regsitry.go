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

package generator

import (
	"fmt"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
)

type registry struct {
	typeHandlers   map[config.DataType]chunks.TypeHandlerFunc
	typeValidators map[config.DataType]chunks.TypeValidatorFunc
}

var Registry = registry{
	typeHandlers:   make(map[config.DataType]chunks.TypeHandlerFunc),
	typeValidators: make(map[config.DataType]chunks.TypeValidatorFunc),
}

func (r *registry) Register(
	dataType config.DataType,
	handlerFunc chunks.TypeHandlerFunc,
	validatorFunc chunks.TypeValidatorFunc) {
	r.typeHandlers[dataType] = handlerFunc
	r.typeValidators[dataType] = validatorFunc
}

func (r *registry) getTypeHandler(dataType config.DataType) (chunks.TypeHandlerFunc, error) {
	if typeHandler, ok := r.typeHandlers[dataType]; ok {
		return typeHandler, nil
	}
	return nil, fmt.Errorf("type '%s' unknown", dataType)
}
func (r *registry) getTypeValidator(dataType config.DataType) (chunks.TypeValidatorFunc, error) {
	if typeValidator, ok := r.typeValidators[dataType]; ok {
		return typeValidator, nil
	}
	return nil, fmt.Errorf("type '%s' unknown", dataType)
}
