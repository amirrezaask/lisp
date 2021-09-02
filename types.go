package main

type NodeType string

const (
	NodeType_Number = "number"
	NodeType_String = "string"
	NodeType_Map    = "map"
	NodeType_List   = "list"
	NodeType_Nil    = "nil"
	NodeType_Atom   = "atom"
)

type Node struct {
	Type  NodeType
	Value interface{}
}

type List struct {
	value []*Node
}

func (l *List) Car() *Node {
	return l.value[0]
}

func (l *List) Cdr() []*Node {
	return l.value[1:]
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
