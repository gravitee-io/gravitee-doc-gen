package schema

import (
	"bytes"
	"encoding/json"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"strconv"
)

type NodeKind int

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

type Object struct {
	baseNode
	Fields map[string]interface{} `json:",inline"`
}

type Array struct {
	baseNode
	Items []interface{} `json:",inline"`
}
type Set map[any]bool

type Attribute struct {
	baseNode
	Value       any `json:",inline"`
	Title       string
	Description string
	Type        string
	When        map[string]Set
	Enums       []any
}

type Value struct {
	Value any `json:",inline"`
}

func NewObject(name string) *Object {
	return &Object{
		baseNode: baseNode{name: name},
		Fields:   make(map[string]interface{}),
	}
}

func NewArray(name string) *Array {
	return &Array{
		baseNode: baseNode{name: name},
		Items:    make([]interface{}, 0),
	}
}

func NewAttribute(name string, value any) *Attribute {
	return &Attribute{
		baseNode: baseNode{name: name},
		Value:    value,
	}
}

func NewValue(value any) Value {
	return Value{
		Value: value,
	}
}

func (o Object) Name() string {
	return o.name
}

func (_ Object) Kind() NodeKind {
	return ObjectNode
}

func (o Object) IsEmpty() bool {
	return len(o.Fields) == 0
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Fields)
}

func (o Object) String() string {
	return o.name + " (object), len: " + strconv.Itoa(len(o.Fields))
}

func (_ Array) Kind() NodeKind {
	return ArrayNode
}

func (a Array) Name() string {
	return a.name
}

func (a Array) IsEmpty() bool {
	return len(a.Items) == 0
}

func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Items)
}

func (a Array) String() string {
	return a.name + " (array), len: " + strconv.Itoa(len(a.Items))
}

func (a Attribute) Name() string {
	return a.name
}

func (_ Attribute) Kind() NodeKind {
	return AttributeNode
}

func (a Attribute) IsEmpty() bool {
	return a.Value == nil
}

func (a Attribute) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	e := json.NewEncoder(&buffer)
	e.SetEscapeHTML(false)
	err := e.Encode(a.Value)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (a Attribute) String() string {
	return a.name + "=" + util.AnyToString(a.Value)
}
func (_ Value) Name() string {
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
