package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/core/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/config"
	"gopkg.in/yaml.v3"
	"os"
)

type ExampleSpecProvider interface {
	ExampleSpecs() []ExampleSpec
	SetConfigData(ConfigData)
	GetConfigData() ConfigData
}

type ConfigData struct {
	GenExamples []GenExampleSpec `yaml:"genExamples"`
	RawExamples []RawExampleSpec `yaml:"rawExamples"`
}

func LoadConfig(chunk config.Chunk, provider ExampleSpecProvider) error {
	examplesFile := chunks.GetString(chunk, "examples")
	bytes, err := os.ReadFile(examplesFile)
	if err != nil {
		return err
	}
	cfg := &ConfigData{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return err
	}

	provider.SetConfigData(*cfg)
	return nil
}
