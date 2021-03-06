package ast

type Node interface {
	Visit(v Visitor) string
}

type Literal struct {
	Value interface{}
}

func (n *Literal) Visit(v Visitor) (s string) {
	s = v.GetLiteral(*n)
	return
}

type And struct {
	Children []Node
}

func (n And) Visit(v Visitor) (s string) {
	s = v.GetAnd(n)
	return
}

type Binary struct {
	Left  Node
	Right Node
}

type Equality struct{ Binary }

func (n Equality) Visit(v Visitor) (s string) {
	s = v.GetEquality(n)
	return
}

type InfixOperation struct{ Binary }

type Join interface {
	Visit(v Visitor) (s string)
	SetRight(n Node)
}

type BaseJoin struct {
	Left  Node
	Right Node
}

type Unary struct {
	Expression Node
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


//-----------------InfixOperation----------------
type Multiplication struct{ InfixOperation }

func (n Multiplication) Visit(v Visitor) (s string) {
	s = v.GetMultiplication(n)
	return
}

type Division struct{ InfixOperation }

func (n Division) Visit(v Visitor) (s string) {
	s = v.GetDivision(n)
	return
}

type Addition struct{ InfixOperation }

func (n Addition) Visit(v Visitor) (s string) {
	s = v.GetAddition(n)
	return
}

type Subtraction struct{ InfixOperation }

func (n Subtraction) Visit(v Visitor) (s string) {
	s = v.GetSubtraction(n)
	return
}


//-----------------Join----------------
type InnerJoin struct{ *BaseJoin }

func (n InnerJoin) Visit(v Visitor) (s string) {
	s = v.GetInnerJoin(n)
	return
}

type OuterJoin struct{ *BaseJoin }

func (n OuterJoin) Visit(v Visitor) (s string) {
	s = v.GetOuterJoin(n)
	return
}

type StringJoin struct{ name string }

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


// ------- NOT GENERATED -----------
func (n *BaseJoin) SetRight(r Node) {
	n.Right = r
}

func (n *BaseJoin) SetLeft(r Node) {
	n.Left = r
}
