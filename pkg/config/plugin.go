package config

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/rickar/props"
)

const PluginChunkId = "Plugin"

func ReadPlugin() (Plugin, error) {

	plugin, err := bootstrap.Registry.UpdateData("plugin", func(data any) (any, error) {
		properties := data.(*props.Properties)
		plugin := Plugin{
			Id:    properties.GetDefault("id", ""),
			Type:  properties.GetDefault("type", ""),
			Title: properties.GetDefault("description", ""),
		}
		return plugin, plugin.Check()
	})

	if err != nil {
		return Plugin{}, err
	}

	return plugin.(Plugin), nil

}

func (p Plugin) Check() error {
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
