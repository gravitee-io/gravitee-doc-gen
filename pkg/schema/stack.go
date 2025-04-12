package schema

type NodeStack struct {
	stack []Node
}

func NewNodeStack(root Node) *NodeStack {
	return &NodeStack{
		stack: []Node{root},
	}
}

func (s *NodeStack) Add(ctx *VisitContext, name string, value interface{}) {
	node := ctx.NodeStack().Peek()
	if node.Type() == ArrayNode {
		array := node.(*Array)
		array.Items = append(array.Items, value)
	} else {
		node.(*Object).Fields[name] = value
	}

	if val, ok := value.(Node); ok {
		ctx.NodeStack().Push(val)
	}
}

func (s *NodeStack) Push(value Node) {
	s.stack = append(s.stack, value)
}

func (s *NodeStack) Peek() Node {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

func (s *NodeStack) Pop() {
	// check if current needs to be removed
	var property string
	var remove bool
	node := s.Peek()
	if node != nil && node.IsEmpty() {
		property = node.Name()
		remove = true
	}

	s.stack = removeLast[Node](s.stack)

	if remove {
		if last := s.Peek(); last.Type() == ArrayNode {
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
	result := make([]string, 0, len(s.stack))
	for _, node := range s.stack {
		result = append(result, node.Name())
	}
	return result
}

func (s *NodeStack) GetAncestorNames() []string {
	ancestors := make([]string, len(s.stack)-1)
	// skip root
	for i := 1; i < len(s.stack); i++ {
		ancestors[i-1] = s.stack[i].Name()
	}
	return ancestors
}
