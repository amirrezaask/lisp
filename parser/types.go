package parser

type NodeType string

const (
	NodeTypeNumber = "number"
	NodeTypeString = "string"
	NodeTypeMap    = "map"
	NodeTypeList   = "list"
	NodeTypeNil    = "nil"
	NodeTypeAtom   = "atom"
)

type Node struct {
	Type  NodeType
	Value interface{}
}

type List struct {
	Value []*Node
}

func (l *List) Car() *Node {
	return l.Value[0]
}

func (l *List) Cdr() []*Node {
	return l.Value[1:]
}

type Number struct {
	value float64
}

type String struct {
	value string
}

type Atom struct {
	value string
}
type Nil struct{}
