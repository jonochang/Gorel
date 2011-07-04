package ast

type Function struct {
	Expressions []Node
	Alias       *SqlLiteral
	Distinct    bool
}


//-----------------Function----------------
type CountNode struct{ Function }

func (n CountNode) Visit(v Visitor) (s string) {
	s = v.GetCount(n)
	return
}

type SumNode struct{ Function }

func (n SumNode) Visit(v Visitor) (s string) {
	s = v.GetSum(n)
	return
}

type Exists struct{ Function }

func (n Exists) Visit(v Visitor) (s string) {
	s = v.GetExists(n)
	return
}

type MaxNode struct{ Function }

func (n MaxNode) Visit(v Visitor) (s string) {
	s = v.GetMax(n)
	return
}

type MinNode struct{ Function }

func (n MinNode) Visit(v Visitor) (s string) {
	s = v.GetMin(n)
	return
}

type Avg struct{ Function }

func (n Avg) Visit(v Visitor) (s string) {
	s = v.GetAvg(n)
	return
}

//---------------- addtional functions ----------
func (function *Function) as(literal string) *Function {
	sql := NewSqlLiteral(literal)
	function.Alias = &sql
	return function
}

func (n CountNode) As(literal string) CountNode { n.as(literal); return n }
func (n SumNode) As(literal string) SumNode     { n.as(literal); return n }
func (n MaxNode) As(literal string) MaxNode     { n.as(literal); return n }
func (n MinNode) As(literal string) MinNode     { n.as(literal); return n }
