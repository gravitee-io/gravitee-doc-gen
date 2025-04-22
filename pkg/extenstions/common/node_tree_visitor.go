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

	if array, ok := node.(*visitor.Array); ok {
		nodeTreeVisitor.OnArrayStart(*array, level)
		for _, node := range array.Items {
			if node.Kind() == visitor.ValueNode {
				nodeTreeVisitor.OnArrayItem(*array, node.(visitor.Value), level)
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
