package common

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"os"
)

func NoopTypeHandler(_ config.Chunk) (chunks.Processed, error) {
	return chunks.Processed{}, nil
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

func GetDataFile(chunk config.Chunk) string {
	return GetFile(chunk, string(chunk.Type))
}

func GetFile(chunk config.Chunk, key string) string {
	return chunk.Data[key].(string)
}
