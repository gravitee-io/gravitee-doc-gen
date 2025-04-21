package filehandlers

import (
	"github.com/rickar/props"
	"os"
)

const PropertiesExt = ".properties"

func PropertiesFileHandler(propFile string) (any, error) {
	file, err := os.Open(propFile)
	if err != nil {
		return nil, err
	}

	properties := props.NewProperties()
	err = properties.Load(file)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
