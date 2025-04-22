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
		return val.(T)
	}
	return defaultValue
}
