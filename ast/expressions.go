package ast

type ExpressionFunctions struct{}

func (e *ExpressionFunctions) count(self Node, distinct bool) CountNode {
	alias := NewSqlLiteral("Count_id")
	return CountNode{Function{[]Node{self}, &alias, distinct}}
}
