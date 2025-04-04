package bootstrap

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Load(rootDir string) error {
	bootstrapFile := filepath.Join(rootDir, "bootstrap.yaml")
	_, err := os.Stat(bootstrapFile)
	if err != nil {
		return err
	}
	type fileContent struct {
		DataFiles []string `yaml:"dataFiles"`
	}
	bootstrap := &fileContent{}
	content, err := os.ReadFile(bootstrapFile)
	err = yaml.Unmarshal(content, bootstrap)
	if err != nil {
		return err
	}

	for _, file := range bootstrap.DataFiles {
		_, err := Registry.load(file)
		if err != nil {
			return err
		}
	}

	return nil

}
