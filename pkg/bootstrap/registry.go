package bootstrap

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"os"
	"path/filepath"
)

type registry struct {
	data     map[string]interface{}
	handlers map[string]Handler
	exports  map[string]string
}

var Registry = &registry{
	data:     make(map[string]interface{}),
	exports:  make(map[string]string),
	handlers: make(map[string]Handler),
}

type Handler func(file string) (any, error)
type PostProcessor func(any) (any, error)

func (r *registry) RegisterHandler(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *registry) UpdateData(name string, processor PostProcessor) (any, error) {
	processed, err := processor(r.GetData(name))
	if err != nil {
		return nil, err
	}
	r.data[name] = processed
	return processed, nil
}

func (r *registry) GetData(name string) any {
	if data, ok := r.data[name]; ok {
		return data
	}
	panic(fmt.Sprintf("'%s' bootstrap data does not exist", name))
}

func (r *registry) GetExports() map[string]string {
	clone := make(map[string]string)
	for k, v := range r.exports {
		clone[k] = v
	}
	return clone
}

func (r *registry) load(file string, export string) (any, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("%s is a directory, should be a file", file)
	}

	if handle, ok := r.handlers[filepath.Ext(file)]; ok {
		val, err := handle(file)
		if err != nil {
			return nil, err
		}
		key := filepath.Base(util.BaseFileNoExt(file))
		r.data[key] = val
		r.exports[key] = export
		return val, nil
	} else {
		panic(fmt.Sprintf("no '%s' handler for bootstrap file: %s ", filepath.Ext(file), file))
	}

}
