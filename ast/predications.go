package ast

type Attribute interface {
	Predications
}

type Predications interface {
	Eq(Literal) Node
}

func (f Field) Eq(l Literal) (n Node) {
	n = Equality{Binary{f, l}}
	return
}
