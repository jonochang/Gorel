package ast

type Visitor interface {
	GetLiteral(n Literal) string
	GetEquality(n Equality) string
	GetAnd(n And) string

	//-----------------And----------------

	//-----------------Binary----------------
	GetBetween(n Between) string
	GetNotEqual(n NotEqual) string
	GetAssignment(n Assignment) string
	GetOr(n Or) string
	GetAs(n As) string
	GetGreaterThan(n GreaterThan) string
	GetGreaterThanOrEqual(n GreaterThanOrEqual) string
	GetLessThan(n LessThan) string
	GetLessThanOrEqual(n LessThanOrEqual) string
	GetMatches(n Matches) string
	GetDoesNotMatch(n DoesNotMatch) string
	GetNotIn(n NotIn) string
	GetOrdering(n Ordering) string
	GetValues(n Values) string
	GetDeleteStatement(n DeleteStatement) string
	GetTableAlias(n TableAlias) string
	GetExcept(n Except) string
	GetIntersect(n Intersect) string
	GetUnion(n Union) string
	GetUnionAll(n UnionAll) string

	//-----------------Equality----------------
	GetIn(n In) string

	//-----------------Function----------------
	GetCount(n CountNode) string
	GetSum(n SumNode) string
	GetExists(n Exists) string
	GetMax(n Max) string
	GetMin(n Min) string
	GetAvg(n Avg) string

	//-----------------InfixOperation----------------
	GetMultiplication(n Multiplication) string
	GetDivision(n Division) string
	GetAddition(n Addition) string
	GetSubtraction(n Subtraction) string

	//-----------------Join----------------
	GetInnerJoin(n InnerJoin) string
	GetOuterJoin(n OuterJoin) string
	GetStringJoin(n StringJoin) string

	//-----------------Unary----------------
	GetNot(n Not) string
	GetLock(n Lock) string
	GetOffset(n Offset) string
	GetLimit(n Limit) string
	GetTop(n Top) string
	GetHaving(n Having) string
	GetUnqualifiedColumn(n UnqualifiedColumn) string
	GetGroup(n Group) string
	GetGrouping(n Grouping) string
	GetOn(n On) string

	//---- not generated ----
	GetField(n Field) string
	GetSelectCore(n SelectCore) string
	GetTable(n Table) string
	GetJoinSource(n JoinSource) string
	GetSqlLiteral(n SqlLiteral) string
	GetSelectStatement(n SelectStatement) string
	GetWith(n With) string
}
