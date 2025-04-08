package bootstrap

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type data struct {
	Filename   string `yaml:"file"`
	ExportedAs string `yaml:"exportedAs"`
}

type fileContent struct {
	Data []data `yaml:"data"`
}

func Load(rootDir string) error {
	bootstrapFile := filepath.Join(rootDir, "bootstrap.yaml")
	_, err := os.Stat(bootstrapFile)
	if err != nil {
		return err
	}

	bootstrap := &fileContent{}
	content, err := os.ReadFile(bootstrapFile)
	err = yaml.Unmarshal(content, bootstrap)
	if err != nil {
		return err
	}

	for _, data := range bootstrap.Data {
		_, err := Registry.load(data.Filename, data.ExportedAs)
		if err != nil {
			return err
		}
	}

	return nil

}
