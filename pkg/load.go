package pkg

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
)

func Load() ([]chunks.Ready, config.Plugin, config.Config, error) {
	pl, err := config.ReadPlugin()
	if err != nil {
		return nil, config.Plugin{}, config.Config{}, err
	}

	cfg, err := config.Read(pl)
	if err != nil {
		return nil, config.Plugin{}, cfg, err
	}

	ready, err := generator.GetReady(cfg.Chunks)
	if err != nil {
		return nil, config.Plugin{}, config.Config{}, err
	}
	return ready, pl, cfg, nil
}
