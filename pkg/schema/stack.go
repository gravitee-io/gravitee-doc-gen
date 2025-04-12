package schema

type NodeStack struct {
	stack []Node
}

func NewNodeStack(root Node) *NodeStack {
	return &NodeStack{
		stack: []Node{root},
	}
}

func (s *NodeStack) add(ctx *VisitContext, toAdd Node) {
	if s == nil {
		panic("stack is nil, cannot add node to stack. stack it should created with the visitor")
	}
	node := ctx.NodeStack().peek()
	if node.Type() == ArrayNode {
		array := node.(*Array)
		array.Items = append(array.Items, toAdd)
	} else if node.Type() == ObjectNode {
		node.(*Object).Fields[toAdd.Name()] = toAdd
	}

	if toAdd.Type() != AttributeNode {
		ctx.NodeStack().push(toAdd)
	}
}

func (s *NodeStack) push(value Node) {
	if s == nil {
		return
	}
	s.stack = append(s.stack, value)
}

func (s *NodeStack) peek() Node {
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
	node := s.peek()
	if node != nil && node.IsEmpty() {
		property = node.Name()
		remove = true
	}

	s.stack = removeLast[Node](s.stack)

	if remove {
		if last := s.peek(); last.Type() == ArrayNode {
			array := last.(*Array)
			array.Items = removeLast(array.Items)
		} else if last.Type() == ObjectNode {
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

func (s *NodeStack) GetNames() []string {
	if s == nil {
		panic("stack is nil, cannot get names. stack it should created with the visitor.")

	}
	result := make([]string, 0, len(s.stack))
	for _, node := range s.stack {
		result = append(result, node.Name())
	}
	return result
}

func (s *NodeStack) GetAncestorNames() []string {
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
