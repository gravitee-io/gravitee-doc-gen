// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package visitor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Two discriminator properties are enough: newOneOfDescriptor used to collect them
// by ranging over a map, so their order flipped between runs and churned the docs.
const twoDiscriminatorSchema = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "oneOf": [
    {
      "title": "Inline json to avro",
      "properties": {
        "schemaLocation": { "const": "inline" },
        "conversion": { "const": "json-to-avro" }
      }
    },
    {
      "title": "Registry avro to json",
      "properties": {
        "schemaLocation": { "const": "schema-registry" },
        "conversion": { "const": "avro-to-json" }
      }
    }
  ]
}`

func compileTestSchema(t *testing.T, body string) *jsonschema.Schema {
	t.Helper()
	file := filepath.Join(t.TempDir(), "schema-form.json")
	if err := os.WriteFile(file, []byte(body), 0o600); err != nil {
		t.Fatalf("write schema: %v", err)
	}
	compiled, err := schema.CompileWithExtensions(file)
	if err != nil {
		t.Fatalf("compile schema: %v", err)
	}
	return compiled
}

func TestNewOneOfDescriptorOrdersSpecsByProperty(t *testing.T) {
	parent := compileTestSchema(t, twoDiscriminatorSchema)

	// Go randomises map iteration per range, so a single pass can pass by luck.
	// 100 passes make an unsorted implementation fail with probability 1-2^-99.
	const passes = 100
	want := []string{"conversion", "schemaLocation"}

	for i := 0; i < passes; i++ {
		got := make([]string, 0, len(want))
		for _, spec := range newOneOfDescriptor(parent).Specs {
			got = append(got, spec.Property)
		}
		if len(got) != len(want) {
			t.Fatalf("pass %d: got %d specs, want %d (%v)", i, len(got), len(want), got)
		}
		for j := range want {
			if got[j] != want[j] {
				t.Fatalf("pass %d: specs out of order: got %v, want %v", i, got, want)
			}
		}
	}
}
