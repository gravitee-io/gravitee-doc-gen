package examples

import (
	"encoding/json"
	"strconv"
)

type Stackable interface {
	IsArray() bool
	Name() string
	IsEmpty() bool
}

type Object struct {
	name   string
	Fields map[string]interface{} `json:",inline"`
}

type Array struct {
	name  string
	Items []interface{} `json:",inline" `
}

func NewObject(name string) *Object {
	return &Object{
		name:   name,
		Fields: make(map[string]interface{}),
	}
}

func NewArray(name string) *Array {
	return &Array{
		name:  name,
		Items: make([]interface{}, 0),
	}
}

func (o Object) Name() string {
	return o.name
}

func (_ Object) IsArray() bool {
	return false
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

func (_ Array) IsArray() bool {
	return true
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
