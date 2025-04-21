package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/core/util"
)

func GenExamplePostProcessor(data any) (any, error) {
	object := data.(*util.Unstructured)
	return util.AnyMapToStruct[GenExamples](object)
}
