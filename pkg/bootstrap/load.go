package bootstrap

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const RootDirDataKey = "rootDir"

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

	// add this here so any one can use it
	registry.data["rootDir"] = rootDir
	registry.exports["rootDir"] = "RootDir"

	content, err := util.RenderTemplateFromFile(bootstrapFile, GetExported())

	bootstrap := &fileContent{}
	err = yaml.Unmarshal(content, bootstrap)
	if err != nil {
		return err
	}

	for _, data := range bootstrap.Data {
		_, err := load(data.Filename, data.ExportedAs)
		if err != nil {
			return err
		}
	}

	return nil

}
