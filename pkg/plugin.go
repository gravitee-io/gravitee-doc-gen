package pkg

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/rickar/props"
	"os"
	"path"
)

type Plugin struct {
	Id    string
	Type  string
	Title string
}

func (p Plugin) String() string {
	return fmt.Sprintf("Plugin{Id: %s, Type: %s, Title: %s}", p.Id, p.Type, p.Title)
}

func PluginPostProcessor(data any) (any, error) {
	properties := data.(*props.Properties)
	plugin := Plugin{
		Id:    properties.GetDefault("id", ""),
		Type:  properties.GetDefault("type", ""),
		Title: properties.GetDefault("description", ""),
	}
	return plugin, plugin.Validate()
}

func (p Plugin) Validate() error {
	if p.Type == "" {
		return errors.New("plugin type is required")
	}
	if p.Title == "" {
		return errors.New("plugin description is required")
	}
	if p.Id == "" {
		return errors.New("plugin id is required")
	}
	return nil
}

func PluginRelatedFile(filename string) (string, error) {
	plugin := bootstrap.GetData("plugin").(Plugin)
	rootDir := bootstrap.GetData(bootstrap.RootDirDataKey).(string)
	specificConfig := path.Join(rootDir, plugin.Type, plugin.Id, filename+".yaml")
	defaultConfig := path.Join(rootDir, plugin.Type, filename)
	if _, err := os.Stat(specificConfig); err == nil {
		return specificConfig, nil
	}
	if _, err := os.Stat(defaultConfig); err == nil {
		return defaultConfig, nil
	}
	return "", errors.New(fmt.Sprintf("plugin related file not found. filename: %s, plugin: %s", filename, plugin))
}
