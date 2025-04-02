package config

import (
	"errors"
	"github.com/rickar/props"
	"os"
)

const file = "src/main/resources/plugin.properties"

const PluginChunkId = "Plugin"

func ReadPlugin() (Plugin, error) {

	file, err := os.Open(file)
	if err != nil {
		return Plugin{}, err
	}
	properties := props.NewProperties()
	err = properties.Load(file)
	if err != nil {
		return Plugin{}, err
	}

	plugin := Plugin{
		Id:    properties.GetDefault("id", ""),
		Type:  properties.GetDefault("type", ""),
		Title: properties.GetDefault("description", ""),
	}

	return plugin, plugin.Check()
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
