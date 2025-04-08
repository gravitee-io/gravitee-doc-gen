package schema

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema/extensions"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func CompileWithExtensions(schemaFile string) (*jsonschema.Schema, error) {

	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true
	compiler.RegisterExtension(extensions.SecretExtension, nil, &extensions.BoolValueCompiler{Ext: extensions.SecretExtension})
	compiler.RegisterExtension(extensions.ElExtension, nil, &extensions.BoolValueCompiler{Ext: extensions.ElExtension})
	compiler.RegisterExtension(extensions.ReadmeExamplesExtension, nil, &extensions.Compiler{})

	return compiler.Compile(schemaFile)

}
