package pkg

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
)

func Load(rootDir string) ([]chunks.Ready, config.Config, error) {

	err := bootstrap.Load(rootDir)
	if err != nil {
		return nil, config.Config{},
			errors.New(fmt.Sprintf("failed to load bootstrap.yaml: %s", err.Error()))
	}

	// TODO remove this
	pl, err := config.ReadPlugin()
	if err != nil {
		return nil, config.Config{}, err
	}

	// TODO pass ConfigFileResolver
	cfg, err := config.Read(rootDir, pl)
	if err != nil {
		return nil, cfg, err
	}

	ready, err := generator.GetReady(cfg.Chunks)
	if err != nil {
		return nil, config.Config{}, err
	}
	return ready, cfg, nil
}
