package core

import (
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/generator"
)

func Load(rootDir string, resolver config.FileResolver) ([]chunks.Generated, config.Config, error) {

	err := bootstrap.Load(rootDir)
	if err != nil {
		return nil, config.Config{},
			errors.New(fmt.Sprintf("failed to load bootstrap.yaml: %s", err.Error()))
	}

	cfg, err := config.Read(rootDir, resolver)
	if err != nil {
		return nil, cfg, err
	}

	ready, err := generator.GetReady(cfg.Chunks)
	if err != nil {
		return nil, config.Config{}, err
	}

	generated, err := generator.Generate(ready)
	if err != nil {
		return nil, config.Config{}, err
	}

	return generated, cfg, nil

}
