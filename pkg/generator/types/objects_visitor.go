package types

import "github.com/gravitee-io-labs/readme-gen/pkg/schema"

type ObjectVisitor interface {
	OnObjectStart(object schema.Object, level int)
	OnObjectEnd(object schema.Object, level int)
	OnArrayStart(array schema.Array, level int)
	OnArrayItem(parent schema.Array, value schema.Value, level int)
	OnArrayEnd(array schema.Array, level int)
	OnAttribute(attribute schema.Attribute, level int)
}

func Visit(ctx *schema.VisitContext, visitor ObjectVisitor) {
	stack := ctx.NodeStack()
	if stack == nil {
		panic("stack is nil")
	}
	stack.Reset()
	root := stack.Peek()
	visit(root, visitor, 0)
}

func visit(node schema.Node, visitor ObjectVisitor, level int) {

	if object, ok := node.(*schema.Object); ok {
		visitor.OnObjectStart(*object, level)
		for _, node := range object.Children() {
			visit(node, visitor, level+1)
		}
		visitor.OnObjectEnd(*object, level)
	}

	if array, ok := node.(*schema.Array); ok {
		visitor.OnArrayStart(*array, level)
		for _, node := range array.Items {
			if node.Kind() == schema.ValueNode {
				visitor.OnArrayItem(*array, node.(schema.Value), level)
			} else {
				visit(node, visitor, level+1)
			}
		}
		visitor.OnArrayEnd(*array, level)
	}

	if attribute, ok := node.(*schema.Attribute); ok {
		visitor.OnAttribute(*attribute, level)
	}

}
