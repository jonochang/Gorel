package ast

type With struct{ Unary }

func (n With) Visit(v Visitor) (s string) {
	s = v.GetWith(n)
	return
}
