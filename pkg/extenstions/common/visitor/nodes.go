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
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type NodeKind int

func (k NodeKind) String() any {
	switch k {
	case ValueNode:
		return "value"
	case ArrayNode:
		return schema.ArrayType
	case ObjectNode:
		return schema.ObjectType
	case AttributeNode:
		return "attribute"
	case Unknown:
		return "unknown"
	}
	panic("unreachable")
}

const (
	Unknown       NodeKind = iota
	ObjectNode    NodeKind = iota
	ArrayNode     NodeKind = iota
	AttributeNode NodeKind = iota
	ValueNode     NodeKind = iota
)

type Node interface {
	Kind() NodeKind
	Name() string
	IsEmpty() bool
}

type baseNode struct {
	name string
}

type documented struct {
	Title       string
	Description string
}

type Object struct {
	baseNode
	documented
	Fields map[string]Node `json:",inline" yaml:",inline"`
	names  []string
	root   bool
}

type Array struct {
	baseNode
	documented
	ItemType string
	Items    []Node `json:",inline" yaml:",inline"`
}

type Attribute struct {
	baseNode
	documented
	parent               *jsonschema.Schema
	IsOneOfProperty      bool
	IsOneOfDiscriminator bool
	Value                any `json:",inline" yaml:",inline"`
	Default              any
	Type                 string
	When                 map[string]util.Set
	Enums                []any
}

type Value struct {
	Value any `json:",inline" yaml:",inline"`
}

func NewObject(name string) *Object {
	return &Object{
		baseNode: baseNode{name: name},
		Fields:   make(map[string]Node),
		names:    make([]string, 0),
	}
}

func NewArray(name string) *Array {
	return &Array{
		baseNode: baseNode{name: name},
		Items:    make([]Node, 0),
	}
}

func NewAttribute(name string, parent *jsonschema.Schema) *Attribute {
	return NewAttributeWithValue(name, nil, parent)
}

func NewAttributeWithValue(name string, value any, parent *jsonschema.Schema) *Attribute {
	return &Attribute{
		baseNode: baseNode{name: name},
		parent:   parent,
		Value:    value,
		When:     make(map[string]util.Set),
	}
}

func NewValue(value any) Value {
	return Value{
		Value: value,
	}
}

func (o *Object) Name() string {
	if o.root {
		return "<root>"
	}
	return o.name
}

func (*Object) Kind() NodeKind {
	return ObjectNode
}

func (o *Object) IsEmpty() bool {
	return len(o.Fields) == 0
}

func (o *Object) AddChild(node Node) {
	o.Fields[node.Name()] = node
	o.names = append(o.names, node.Name())
}

func (o *Object) Children() []Node {
	nodes := make([]Node, 0)
	for _, name := range o.names {
		nodes = append(nodes, o.Fields[name])
	}
	return nodes
}

func (o *Object) IsDiscriminator(property string) bool {
	if attribute, ok := o.Fields[property]; ok {
		return util.As[*Attribute](attribute).IsOneOfDiscriminator
	}
	return false
}

func (o *Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Fields)
}

func (o *Object) String() string {
	return o.name + " (object), len: " + strconv.Itoa(len(o.Fields))
}

func (*Array) Kind() NodeKind {
	return ArrayNode
}

func (a *Array) Name() string {
	return a.name
}

func (a *Array) IsEmpty() bool {
	return len(a.Items) == 0
}

func (a *Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Items)
}

func (a *Array) String() string {
	return a.name + " (array), len: " + strconv.Itoa(len(a.Items))
}

func (a *Array) Values() []any {
	values := make([]any, 0)
	for _, item := range a.Items {
		if val, ok := item.(Value); ok {
			values = append(values, val.Value)
		}
	}
	return values
}

func (a *Attribute) Name() string {
	return a.name
}

func (*Attribute) Kind() NodeKind {
	return AttributeNode
}

func (a *Attribute) IsEmpty() bool {
	return a.Value == nil
}

func (a *Attribute) updateWhen(ctx *VisitContext) {
	for _, spec := range ctx.PeekOneOf().Specs {
		list := NewSchemaPropertyList(a.parent)
		value := GetValueOrFirstExample(list.Get(spec.Property), ctx)
		var set util.Set
		if s, ok := a.When[spec.Property]; ok {
			set = s
		} else {
			set = util.NewSet()
		}
		set.Add(value)
		a.When[spec.Property] = set
	}
}

func (a *Attribute) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	e := json.NewEncoder(&buffer)
	e.SetEscapeHTML(false)
	err := e.Encode(a.Value)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (a *Attribute) String() string {
	return a.name + "=" + util.AnyToString(a.Value)
}
func (Value) Name() string {
	return ""
}

func (v Value) Kind() NodeKind {
	return ValueNode
}

func (v Value) IsEmpty() bool {
	return v.Value == nil
}

func (v Value) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	e := json.NewEncoder(&buffer)
	e.SetEscapeHTML(false)
	err := e.Encode(v.Value)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (v Value) String() string {
	return util.AnyToString(v.Value)
}
