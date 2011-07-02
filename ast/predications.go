package ast

type Attribute interface {
	Predications
}

type Predications interface {
	Eq(Node) Node
}

func (f Field) Eq(l Node) (n Node) {
	n = Equality{Binary{f, l}}
	return
}
