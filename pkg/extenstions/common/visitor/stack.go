package visitor

type NodeStack struct {
	stack []Node
}

func NewNodeStack(root *Object) *NodeStack {
	root.root = true
	return &NodeStack{
		stack: []Node{root},
	}
}

func (s *NodeStack) Reset() {
	for {
		if object, ok := s.Peek().(*Object); ok && object.root {
			break
		}
		s.pop()
	}
}

func (s *NodeStack) Nodes() []Node {
	clone := make([]Node, len(s.stack))
	copy(clone, s.stack)
	return clone
}

func (s *NodeStack) Properties() []string {
	if s == nil {
		panic("stack is nil, cannot get names. stack it should created with the visitor.")

	}
	result := make([]string, 0, len(s.stack))
	for _, node := range s.stack {
		result = append(result, node.Name())
	}
	return result
}

func (s *NodeStack) GetAncestorProperty() []string {
	if s == nil {
		panic("stack is nil, cannot get ancestors. stack it should created with the visitor.")
	}
	ancestors := make([]string, len(s.stack)-1)
	// skip root
	for i := 1; i < len(s.stack); i++ {
		ancestors[i-1] = s.stack[i].Name()
	}
	return ancestors
}

func (s *NodeStack) add(ctx *VisitContext, added Node) {
	if s == nil {
		panic("stack is nil, cannot add node to stack. stack it should created with the visitor")
	}
	node := ctx.NodeStack().Peek()
	if node.Kind() == ArrayNode {
		array := node.(*Array)
		array.Items = append(array.Items, added)
	} else if node.Kind() == ObjectNode {
		node.(*Object).AddChild(added)
	}
	if added.Kind() == AttributeNode {
		attribute := added.(*Attribute)
		if attribute.IsOneOfProperty {
			attribute.updateWhen(ctx)
		}
	}

	if added.Kind() == ObjectNode || added.Kind() == ArrayNode {
		ctx.NodeStack().push(added)
	}
}

func (s *NodeStack) push(value Node) {
	if s == nil {
		return
	}
	s.stack = append(s.stack, value)
}

func (s *NodeStack) Peek() Node {
	if s == nil {
		panic("stack is nil, cannot add node to stack. stack it should created with the visitor.")
	}
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

func (s *NodeStack) pop() {
	if s == nil {
		return
	}
	// check if current needs to be removed
	var property string
	var remove bool
	node := s.Peek()
	if node != nil && node.IsEmpty() {
		property = node.Name()
		remove = true
	}

	s.stack = removeLast[Node](s.stack)

	if remove && !isRoot(node) {
		if last := s.Peek(); last.Kind() == ArrayNode {
			array := last.(*Array)
			array.Items = removeLast(array.Items)
		} else if last.Kind() == ObjectNode {
			object := last.(*Object)
			delete(object.Fields, property)
		}
	}

}

func removeLast[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

func isRoot(n Node) bool {
	if n.Kind() == ObjectNode {
		return n.(*Object).root
	}
	return false
}
