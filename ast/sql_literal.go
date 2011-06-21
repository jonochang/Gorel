package ast

type SqlLiteral struct {
	Value string
}

func (n SqlLiteral) Visit(v Visitor) (s string) {
	s = v.GetSqlLiteral(n)
	return
}
