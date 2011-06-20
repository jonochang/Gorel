package ast

type SelectCore struct {
	Top           Node
	Projections   []Node
	Wheres        []Node
	Groups        []Node
	Having        Node
	Source        Node
	SetQuantifier Node
}

func (n SelectCore) Visit(v Visitor) (s string) {
	s = v.GetSelectCore(n)
	return
}
