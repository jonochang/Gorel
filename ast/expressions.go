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

func (e *ExpressionFunctions) max(self Node) MaxNode {
	alias := NewSqlLiteral("Max_id")
	return MaxNode{Function{[]Node{self}, &alias, false}}
}

func (e *ExpressionFunctions) min(self Node) MinNode {
	alias := NewSqlLiteral("Min_id")
	return MinNode{Function{[]Node{self}, &alias, false}}
}

func (e *ExpressionFunctions) avg(self Node) AvgNode {
	alias := NewSqlLiteral("Avg_id")
	return AvgNode{Function{[]Node{self}, &alias, false}}
}
