package generator

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/config"
)

type registry struct {
	typeHandlers   map[config.DataType]chunks.TypeHandlerFunc
	typeValidators map[config.DataType]chunks.TypeValidatorFunc
}

var Registry = registry{
	typeHandlers:   make(map[config.DataType]chunks.TypeHandlerFunc),
	typeValidators: make(map[config.DataType]chunks.TypeValidatorFunc),
}

func (r *registry) Register(dataType config.DataType, handlerFunc chunks.TypeHandlerFunc, validatorFunc chunks.TypeValidatorFunc) {
	r.typeHandlers[dataType] = handlerFunc
	r.typeValidators[dataType] = validatorFunc
}

func (r *registry) getTypeHandler(dataType config.DataType) (chunks.TypeHandlerFunc, error) {
	if typeHandler, ok := r.typeHandlers[dataType]; ok {
		return typeHandler, nil
	}
	return nil, fmt.Errorf("type '%s' unknown", dataType)
}
func (r *registry) getTypeValidator(dataType config.DataType) (chunks.TypeValidatorFunc, error) {
	if typeValidator, ok := r.typeValidators[dataType]; ok {
		return typeValidator, nil
	}
	return nil, fmt.Errorf("type '%s' unknown", dataType)
}
