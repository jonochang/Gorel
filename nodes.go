package gorel

type Node interface {
	Visit(v Visitor) string
}

type Literal struct {
	value interface{}
}

func (n Literal) Visit(v Visitor) (s string) {
	s = v.GetLiteral(n)
	return
}

type And struct {
	children []Node
}

type Binary struct {
	left  Node
	right Node
}

type Equality struct{ Binary }

func (n Equality) Visit(v Visitor) (s string) {
	s = v.GetEquality(n)
	return
}

type Function struct {
	expressions []Node
	alias       Literal
	distinct    bool
}

type Join struct{ Binary }

type Unary struct {
	expression Node
}

//-----------------And----------------

//-----------------Binary----------------
type Between struct{ Binary }

func (n Between) Visit(v Visitor) (s string) {
	s = v.GetBetween(n)
	return
}

type NotEqual struct{ Binary }

func (n NotEqual) Visit(v Visitor) (s string) {
	s = v.GetNotEqual(n)
	return
}

type Assignment struct{ Binary }

func (n Assignment) Visit(v Visitor) (s string) {
	s = v.GetAssignment(n)
	return
}

type Or struct{ Binary }

func (n Or) Visit(v Visitor) (s string) {
	s = v.GetOr(n)
	return
}

type As struct{ Binary }

func (n As) Visit(v Visitor) (s string) {
	s = v.GetAs(n)
	return
}

type GreaterThan struct{ Binary }

func (n GreaterThan) Visit(v Visitor) (s string) {
	s = v.GetGreaterThan(n)
	return
}

type GreaterThanOrEqual struct{ Binary }

func (n GreaterThanOrEqual) Visit(v Visitor) (s string) {
	s = v.GetGreaterThanOrEqual(n)
	return
}

type LessThan struct{ Binary }

func (n LessThan) Visit(v Visitor) (s string) {
	s = v.GetLessThan(n)
	return
}

type LessThanOrEqual struct{ Binary }

func (n LessThanOrEqual) Visit(v Visitor) (s string) {
	s = v.GetLessThanOrEqual(n)
	return
}

type Matches struct{ Binary }

func (n Matches) Visit(v Visitor) (s string) {
	s = v.GetMatches(n)
	return
}

type DoesNotMatch struct{ Binary }

func (n DoesNotMatch) Visit(v Visitor) (s string) {
	s = v.GetDoesNotMatch(n)
	return
}

type NotIn struct{ Binary }

func (n NotIn) Visit(v Visitor) (s string) {
	s = v.GetNotIn(n)
	return
}

type Ordering struct{ Binary }

func (n Ordering) Visit(v Visitor) (s string) {
	s = v.GetOrdering(n)
	return
}

type Values struct{ Binary }

func (n Values) Visit(v Visitor) (s string) {
	s = v.GetValues(n)
	return
}

type DeleteStatement struct{ Binary }

func (n DeleteStatement) Visit(v Visitor) (s string) {
	s = v.GetDeleteStatement(n)
	return
}

type TableAlias struct{ Binary }

func (n TableAlias) Visit(v Visitor) (s string) {
	s = v.GetTableAlias(n)
	return
}

type Except struct{ Binary }

func (n Except) Visit(v Visitor) (s string) {
	s = v.GetExcept(n)
	return
}

type Intersect struct{ Binary }

func (n Intersect) Visit(v Visitor) (s string) {
	s = v.GetIntersect(n)
	return
}

type Union struct{ Binary }

func (n Union) Visit(v Visitor) (s string) {
	s = v.GetUnion(n)
	return
}

type UnionAll struct{ Binary }

func (n UnionAll) Visit(v Visitor) (s string) {
	s = v.GetUnionAll(n)
	return
}


//-----------------Equality----------------
type In struct{ Equality }

func (n In) Visit(v Visitor) (s string) {
	s = v.GetIn(n)
	return
}


//-----------------Function----------------
type Count struct{ Function }

func (n Count) Visit(v Visitor) (s string) {
	s = v.GetCount(n)
	return
}

type Sum struct{ Function }

func (n Sum) Visit(v Visitor) (s string) {
	s = v.GetSum(n)
	return
}

type Exists struct{ Function }

func (n Exists) Visit(v Visitor) (s string) {
	s = v.GetExists(n)
	return
}

type Max struct{ Function }

func (n Max) Visit(v Visitor) (s string) {
	s = v.GetMax(n)
	return
}

type Min struct{ Function }

func (n Min) Visit(v Visitor) (s string) {
	s = v.GetMin(n)
	return
}

type Avg struct{ Function }

func (n Avg) Visit(v Visitor) (s string) {
	s = v.GetAvg(n)
	return
}


//-----------------Join----------------
type InnerJoin struct{ Join }

func (n InnerJoin) Visit(v Visitor) (s string) {
	s = v.GetInnerJoin(n)
	return
}

type OuterJoin struct{ Join }

func (n OuterJoin) Visit(v Visitor) (s string) {
	s = v.GetOuterJoin(n)
	return
}

type StringJoin struct{ Join }

func (n StringJoin) Visit(v Visitor) (s string) {
	s = v.GetStringJoin(n)
	return
}


//-----------------Unary----------------
type Not struct{ Unary }

func (n Not) Visit(v Visitor) (s string) {
	s = v.GetNot(n)
	return
}

type Lock struct{ Unary }

func (n Lock) Visit(v Visitor) (s string) {
	s = v.GetLock(n)
	return
}

type Offset struct{ Unary }

func (n Offset) Visit(v Visitor) (s string) {
	s = v.GetOffset(n)
	return
}

type Limit struct{ Unary }

func (n Limit) Visit(v Visitor) (s string) {
	s = v.GetLimit(n)
	return
}

type Top struct{ Unary }

func (n Top) Visit(v Visitor) (s string) {
	s = v.GetTop(n)
	return
}

type Having struct{ Unary }

func (n Having) Visit(v Visitor) (s string) {
	s = v.GetHaving(n)
	return
}

type UnqualifiedColumn struct{ Unary }

func (n UnqualifiedColumn) Visit(v Visitor) (s string) {
	s = v.GetUnqualifiedColumn(n)
	return
}

type Group struct{ Unary }

func (n Group) Visit(v Visitor) (s string) {
	s = v.GetGroup(n)
	return
}

type Grouping struct{ Unary }

func (n Grouping) Visit(v Visitor) (s string) {
	s = v.GetGrouping(n)
	return
}

type On struct{ Unary }

func (n On) Visit(v Visitor) (s string) {
	s = v.GetOn(n)
	return
}
