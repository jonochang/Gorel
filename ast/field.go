package ast

type Field struct {
	Name string
}

func (n Field) Visit(v Visitor) (s string) {
	s = v.GetField(n)
	return
}
