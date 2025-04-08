package pkg

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
	"os"
	"path"
)

func Load(rootDir string) ([]chunks.Ready, config.Config, error) {

	err := bootstrap.Load(rootDir)
	if err != nil {
		return nil, config.Config{},
			errors.New(fmt.Sprintf("failed to load bootstrap.yaml: %s", err.Error()))
	}

	// TODO remove this
	_, err = config.ReadPlugin()
	if err != nil {
		return nil, config.Config{}, err
	}

	cfg, err := config.Read(rootDir, resolveConfigFile)
	if err != nil {
		return nil, cfg, err
	}

	ready, err := generator.GetReady(cfg.Chunks)
	if err != nil {
		return nil, config.Config{}, err
	}
	return ready, cfg, nil
}

func resolveConfigFile(rootDir string) (string, error) {
	data := bootstrap.Registry.GetData("plugin")
	plugin := data.(config.Plugin)
	specificConfig := path.Join(rootDir, plugin.Type, plugin.Id+".yaml")
	defaultConfig := path.Join(rootDir, plugin.Type, "/default.yaml")
	var configFile string
	if stat, err := os.Stat(specificConfig); err == nil && !stat.IsDir() {
		configFile = specificConfig
	} else if stat, err := os.Stat(defaultConfig); err == nil && !stat.IsDir() {
		configFile = defaultConfig
	} else {
		return "", errors.New(fmt.Sprintf("Cannot find %s or %s ", specificConfig, defaultConfig))
	}
	return configFile, nil
}
