package ast

type ExpressionFunctions struct{}

func (e *ExpressionFunctions) count(self Node, distinct bool) CountNode {
	alias := NewSqlLiteral("Count_id")
	return CountNode{Function{[]Node{self}, &alias, distinct}}
}

func (e *ExpressionFunctions) sum(self Node) SumNode {
	alias := NewSqlLiteral("Sum_id")
	return SumNode{Function{[]Node{self}, &alias, false}}
}
