package schema

import (
	"encoding/json"
	"strconv"
)

type NodeType int

const (
	AttributeNode          = iota
	ObjectNode    NodeType = 1
	ArrayNode     NodeType = 2
)

type Node interface {
	Type() NodeType
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
	Items []interface{} `json:",inline" `
}

type Attribute struct {
	baseNode
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

func NewAttribute(name string) *Attribute {
	return &Attribute{
		baseNode: baseNode{name: name},
	}
}

func (a Attribute) Name() string {
	return a.name
}

func (_ Attribute) Type() NodeType {
	return ObjectNode
}

func (a Attribute) IsEmpty() bool {
	return a.Value == nil
}

func (o Object) Name() string {
	return o.name
}

func (_ Object) Type() NodeType {
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

func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Items)
}

func (_ Array) Type() NodeType {
	return ArrayNode
}

func (a Array) Name() string {
	return a.name
}

func (a Array) IsEmpty() bool {
	return len(a.Items) == 0
}

func (a Array) String() string {
	return a.name + " (array), len: " + strconv.Itoa(len(a.Items))
}
