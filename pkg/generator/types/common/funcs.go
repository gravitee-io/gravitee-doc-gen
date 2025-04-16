package common

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"os"
)

func NoopTypeHandler(_ config.Chunk) (chunks.Processed, error) {
	return chunks.Processed{
		Data: bootstrap.GetExported(),
	}, nil
}

func TemplateExistsTypeValidator(chunk config.Chunk) (bool, error) {
	var err error
	exists := FileExists(chunk.Template)
	if chunk.Required && !exists {
		err = errors.New(chunk.Template + " is required but file is missing")
	}
	return exists, err
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
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
