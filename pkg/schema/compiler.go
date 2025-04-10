package schema

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema/extensions"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func CompileWithExtensions(schemaFile string) (*jsonschema.Schema, error) {
	return CompilerWithExtensions().Compile(schemaFile)

}

func CompilerWithExtensions() *jsonschema.Compiler {
	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true
	compiler.RegisterExtension(extensions.GioConfigExtension, nil, &extensions.GioConfigCompiler{})
	return compiler
}
