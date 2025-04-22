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

package common

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
)

type NodeTreeVisitor interface {
	OnObjectStart(object visitor.Object, level int)
	OnObjectEnd(object visitor.Object, level int)
	OnArrayStart(array visitor.Array, level int)
	OnArrayItem(parent visitor.Array, value visitor.Value, level int)
	OnArrayEnd(array visitor.Array, level int)
	OnAttribute(attribute visitor.Attribute, level int)
}

func Visit(stack *visitor.NodeStack, visitor NodeTreeVisitor) {
	if stack == nil {
		panic("stack is nil")
	}
	stack.Reset()
	root := stack.Peek()
	visit(root, visitor, 0)
}

func visit(node visitor.Node, nodeTreeVisitor NodeTreeVisitor, level int) {
	if object, ok := node.(*visitor.Object); ok {
		nodeTreeVisitor.OnObjectStart(*object, level)
		for _, node := range object.Children() {
			visit(node, nodeTreeVisitor, level+1)
		}
		nodeTreeVisitor.OnObjectEnd(*object, level)
	}

	if array, isArray := node.(*visitor.Array); isArray {
		nodeTreeVisitor.OnArrayStart(*array, level)
		for _, node := range array.Items {
			if node.Kind() == visitor.ValueNode {
				if value, isValue := node.(visitor.Value); isValue {
					nodeTreeVisitor.OnArrayItem(*array, value, level)
				}
			} else {
				visit(node, nodeTreeVisitor, level+1)
			}
		}
		nodeTreeVisitor.OnArrayEnd(*array, level)
	}

	if attribute, ok := node.(*visitor.Attribute); ok {
		nodeTreeVisitor.OnAttribute(*attribute, level)
	}
}
