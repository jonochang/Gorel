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
	js := JoinSource{nil, make([]Node, 0)}
	n.Cores = []Node{&SelectCore{Source: &js}}
	return
}
