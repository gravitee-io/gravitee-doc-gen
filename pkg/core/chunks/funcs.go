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

package chunks

import (
	"errors"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

type TypeHandlerFunc func(configChunk config.Chunk) (Processed, error)
type TypeValidatorFunc func(configChunk config.Chunk) (bool, error)

func NoopTypeHandler(_ config.Chunk) (Processed, error) {
	return Processed{
		Data: bootstrap.GetExported(),
	}, nil
}

func TemplateExistsTypeValidator(chunk config.Chunk) (bool, error) {
	var err error
	exists := util.FileExists(chunk.Template)
	if chunk.Required && !exists {
		err = errors.New(chunk.Template + " is required but file is missing")
	}
	return exists, err
}

func GetDataTypeFile(chunk config.Chunk) string {
	return GetString(chunk, string(chunk.Type))
}

func GetString(chunk config.Chunk, key string) string {
	return util.AnyToString(chunk.Data[key])
}

func GetDataOrDefault[T any](chunk config.Chunk, key string, defaultValue T) T {
	if val := chunk.Data[key]; val != nil {
		if t, ok := val.(T); ok {
			return t
		}
	}
	return defaultValue
}
