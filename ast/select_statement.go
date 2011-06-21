package ast

type SelectStatement struct {
	Cores  []Node
	Orders []Node
	Limit  Node
	Offset Node
	Lock   Node
	With   Node
}

func (n SelectStatement) Visit(v Visitor) (s string) {
	s = v.GetSelectStatement(n)
	return
}

func NewSelectStatement() (n SelectStatement) {
	n.Cores = []Node{SelectCore{}}
	return
}
