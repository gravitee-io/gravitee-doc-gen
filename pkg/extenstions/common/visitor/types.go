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
	"sort"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(ctx *VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *Attribute
	OnObjectStart(ctx *VisitContext, property string, object *jsonschema.Schema) *Object
	OnObjectEnd(ctx *VisitContext)
	OnArrayStart(ctx *VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) (*Array, []Value)
	OnArrayEnd(ctx *VisitContext, itemTypeIsObject bool)
	OnOneOfStart(ctx *VisitContext, oneOf *jsonschema.Schema)
	OnOneOf(visitCtx *VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema)
	OnOneOfEnd(*VisitContext)
}

type DiscriminatorSpec struct {
	Values   []any
	Property string
	Type     string
}

type SchemaProperty struct {
	name   string
	schema *jsonschema.Schema
}

type SchemaPropertyList []SchemaProperty

func (l *SchemaPropertyList) Get(name string) *jsonschema.Schema {
	for _, property := range *l {
		if property.name == name {
			return property.schema
		}
	}
	return nil
}

func (l *SchemaPropertyList) Add(name string, attribute *jsonschema.Schema) {
	*l = append(*l, SchemaProperty{name, attribute})
}

func (l *SchemaPropertyList) Sort() {
	list := *l
	sort.Slice(list, func(i, j int) bool {
		return (list)[i].name < (list)[j].name
	})
	*l = list
}
