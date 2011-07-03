package ast

type SqlLiteral struct {
	Value string
	*ExpressionFunctions
}

func (n SqlLiteral) Visit(v Visitor) (s string) {
	s = v.GetSqlLiteral(n)
	return
}

func NewSqlLiteral(value string) SqlLiteral {
	return SqlLiteral{value, &ExpressionFunctions{}}
}

func (s SqlLiteral) Count() CountNode { return s.count(s, false) }

func (s SqlLiteral) CountDistinct() CountNode { return s.count(s, true) }
