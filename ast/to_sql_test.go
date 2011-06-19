package gorel

import (
	"testing"
)

func TestGetLiteral(t *testing.T) {
	v := new(ToSql)
	l := new(Literal)

	t.Log("Test string")
	l.value = "test"
	s := v.GetLiteral(*l)
	if s != "\"test\"" {
		t.Errorf("failed test string")
	}

	t.Log("Test boolean")
	l.value = false
	s = v.GetLiteral(*l)
	if s != "false" {
		t.Errorf("failed test boolean")
	}

	t.Log("Test int")
	l.value = 123
	s = v.GetLiteral(*l)
	if s != "123" {
		t.Errorf("failed test int")
	}

	t.Log("Test float")
	l.value = 123.3
	s = v.GetLiteral(*l)
	if s != "123.3" {
		t.Errorf("failed test float")
	}

	t.Log("Test Array of ints")
	l.value = []int{1, 2, 3, 4}
	s = v.GetLiteral(*l)
	if s != "1,2,3,4" {
		t.Log(s)
		t.Errorf("failed test float")
	}

	t.Log("Test Array of ints")
	l.value = []int{1, 2, 3, 4}
	s = v.GetLiteral(*l)
	if s != "1,2,3,4" {
		t.Log(s)
		t.Errorf("failed test float")
	}

	t.Log("Test Array of strings")
	l.value = []string{"a", "bc", "d", "ef1"}
	s = v.GetLiteral(*l)
	if s != "\"a\",\"bc\",\"d\",\"ef1\"" {
		t.Log(s)
		t.Errorf("failed test float")
	}
}

func TestVisitNilNodes(t *testing.T) {
	v := new(ToSql)
	s := v.VisitNodes(nil)

	if len(s) != 0 {
		t.Errorf("failed to visit nil nodes")
	}

	nodes := make([]Node, 3)
	s = v.VisitNodes(nodes)
	if len(s) != 0 {
		t.Errorf("failed to skip nil nodes")
	}
}

//-----------------And----------------
func TestGetAnd(t *testing.T) {
	v := new(ToSql)
	n := new(And)

	n1 := Literal{1}
	n2 := Literal{2}
	n3 := Literal{3}
	n.children = []Node{n1, n2, n3}

	s := v.GetAnd(*n)
	if s != "1 AND 2 AND 3" {
		t.Log(s)
		t.Errorf("failed to get And")
	}
}

//-----------------Binary----------------
func TestGetBetween(t *testing.T) {
	v := new(ToSql)
	n := new(Between)
	n.left = Literal{1}
	n.right = And{[]Node{
		Literal{0},
		Literal{3}}}

	s := v.GetBetween(*n)
	if s != "1 BETWEEN 0 AND 3" {
		t.Log(s)
		t.Errorf("failed to get Between ")
	}
}

func TestGetNotEqual(t *testing.T) {
	v := new(ToSql)
	n := new(NotEqual)
	n.left = Literal{1}

	s := v.GetNotEqual(*n)
	if s != "1 IS NOT NULL" {
		t.Log(s)
		t.Errorf("failed to get NotEqual null right")
	}

	n.right = Literal{2}
	s = v.GetNotEqual(*n)
	if s != "1 != 2" {
		t.Log(s)
		t.Errorf("failed to get NotEqual 1 != 2")
	}
}

func TestGetAssignment(t *testing.T) {
	v := new(ToSql)
	n := new(Assignment)

	s := v.GetAssignment(*n)
	if s != "" {
		t.Errorf("failed to get Assignment ")
	}
}

func TestGetOr(t *testing.T) {
	v := new(ToSql)
	n := new(Or)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetOr(*n)
	if s != "1 OR 2" {
		t.Log(s)
		t.Errorf("failed to get Or ")
	}
}

func TestGetAs(t *testing.T) {
	v := new(ToSql)
	n := new(As)

	s := v.GetAs(*n)
	if s != "" {
		t.Errorf("failed to get As ")
	}
}

func TestGetGreaterThan(t *testing.T) {
	v := new(ToSql)
	n := new(GreaterThan)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetGreaterThan(*n)
	if s != "1 > 2" {
		t.Log(s)
		t.Errorf("failed to get GreaterThan ")
	}
}

func TestGetGreaterThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(GreaterThanOrEqual)

	n.left = Literal{2}
	n.right = Literal{1}

	s := v.GetGreaterThanOrEqual(*n)
	if s != "2 >= 1" {
		t.Log(s)
		t.Errorf("failed to get GreaterThanOrEqual ")
	}
}

func TestGetLessThan(t *testing.T) {
	v := new(ToSql)
	n := new(LessThan)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetLessThan(*n)
	if s != "1 < 2" {
		t.Log(s)
		t.Errorf("failed to get LessThan ")
	}
}

func TestGetLessThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(LessThanOrEqual)

	n.left = Literal{2}
	n.right = Literal{1}

	s := v.GetLessThanOrEqual(*n)
	if s != "2 <= 1" {
		t.Log(s)
		t.Errorf("failed to get LessThanOrEqual ")
	}
}

func TestGetMatches(t *testing.T) {
	v := new(ToSql)
	n := new(Matches)

	n.left = Literal{2}
	n.right = Literal{1}

	s := v.GetMatches(*n)
	if s != "2 LIKE 1" {
		t.Log(s)
		t.Errorf("failed to get Matches ")
	}
}

func TestGetDoesNotMatch(t *testing.T) {
	v := new(ToSql)
	n := new(DoesNotMatch)

	n.left = Literal{2}
	n.right = Literal{1}

	s := v.GetDoesNotMatch(*n)
	if s != "2 NOT LIKE 1" {
		t.Log(s)
		t.Errorf("failed to get DoesNotMatch ")
	}
}

func TestGetNotIn(t *testing.T) {
	v := new(ToSql)
	n := new(NotIn)

	n.left = Literal{"abc"}
	n.right = Literal{[]string{"1", "2", "3", "4"}}
	s := v.GetNotIn(*n)
	if s != "\"abc\" NOT IN (\"1\",\"2\",\"3\",\"4\")" {
		t.Log(s)
		t.Errorf("failed to get NotIn ")
	}
}

func TestGetOrdering(t *testing.T) {
	v := new(ToSql)
	n := new(Ordering)

	s := v.GetOrdering(*n)
	if s != "" {
		t.Errorf("failed to get Ordering ")
	}
}
func TestGetValues(t *testing.T) {
	v := new(ToSql)
	n := new(Values)

	s := v.GetValues(*n)
	if s != "" {
		t.Errorf("failed to get Values ")
	}
}
func TestGetDeleteStatement(t *testing.T) {
	v := new(ToSql)
	n := new(DeleteStatement)

	s := v.GetDeleteStatement(*n)
	if s != "" {
		t.Errorf("failed to get DeleteStatement ")
	}
}
func TestGetTableAlias(t *testing.T) {
	v := new(ToSql)
	n := new(TableAlias)

	s := v.GetTableAlias(*n)
	if s != "" {
		t.Errorf("failed to get TableAlias ")
	}
}
func TestGetExcept(t *testing.T) {
	v := new(ToSql)
	n := new(Except)

	s := v.GetExcept(*n)
	if s != "" {
		t.Errorf("failed to get Except ")
	}
}
func TestGetIntersect(t *testing.T) {
	v := new(ToSql)
	n := new(Intersect)

	s := v.GetIntersect(*n)
	if s != "" {
		t.Errorf("failed to get Intersect ")
	}
}
func TestGetUnion(t *testing.T) {
	v := new(ToSql)
	n := new(Union)

	s := v.GetUnion(*n)
	if s != "" {
		t.Errorf("failed to get Union ")
	}
}
func TestGetUnionAll(t *testing.T) {
	v := new(ToSql)
	n := new(UnionAll)

	s := v.GetUnionAll(*n)
	if s != "" {
		t.Errorf("failed to get UnionAll ")
	}
}

//-----------------Equality----------------
func TestGetEquality(t *testing.T) {
	v := new(ToSql)
	n := new(Equality)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetEquality(*n)
	if s != "1 = 2" {
		t.Log(s)
		t.Errorf("failed to get Equality")
	}
}

func TestGetIn(t *testing.T) {
	v := new(ToSql)
	n := new(In)

	n.left = Literal{"abc"}
	n.right = Literal{[]string{"1", "2", "3", "4"}}

	s := v.GetIn(*n)
	if s != "\"abc\" IN (\"1\",\"2\",\"3\",\"4\")" {
		t.Log(s)
		t.Errorf("failed to get In ")
	}
}

//-----------------Function----------------
func TestGetCount(t *testing.T) {
	v := new(ToSql)
	n := new(Count)

	s := v.GetCount(*n)
	if s != "" {
		t.Errorf("failed to get Count ")
	}
}
func TestGetSum(t *testing.T) {
	v := new(ToSql)
	n := new(Sum)

	s := v.GetSum(*n)
	if s != "" {
		t.Errorf("failed to get Sum ")
	}
}
func TestGetExists(t *testing.T) {
	v := new(ToSql)
	n := new(Exists)

	s := v.GetExists(*n)
	if s != "" {
		t.Errorf("failed to get Exists ")
	}
}
func TestGetMax(t *testing.T) {
	v := new(ToSql)
	n := new(Max)

	s := v.GetMax(*n)
	if s != "" {
		t.Errorf("failed to get Max ")
	}
}
func TestGetMin(t *testing.T) {
	v := new(ToSql)
	n := new(Min)

	s := v.GetMin(*n)
	if s != "" {
		t.Errorf("failed to get Min ")
	}
}
func TestGetAvg(t *testing.T) {
	v := new(ToSql)
	n := new(Avg)

	s := v.GetAvg(*n)
	if s != "" {
		t.Errorf("failed to get Avg ")
	}
}

//-----------------InfixOperation----------------
func TestGetMultiplication(t *testing.T) {
	v := new(ToSql)
	n := new(Multiplication)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetMultiplication(*n)
	if s != "1 * 2" {
		t.Log(s)
		t.Errorf("failed to get Multiplication ")
	}
}

func TestGetDivision(t *testing.T) {
	v := new(ToSql)
	n := new(Division)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetDivision(*n)
	if s != "1 / 2" {
		t.Log(s)
		t.Errorf("failed to get Division ")
	}
}

func TestGetAddition(t *testing.T) {
	v := new(ToSql)
	n := new(Addition)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetAddition(*n)
	if s != "1 + 2" {
		t.Log(s)
		t.Errorf("failed to get Addition ")
	}
}

func TestGetSubtraction(t *testing.T) {
	v := new(ToSql)
	n := new(Subtraction)

	n.left = Literal{1}
	n.right = Literal{2}

	s := v.GetSubtraction(*n)
	if s != "1 - 2" {
		t.Log(s)
		t.Errorf("failed to get Subtraction ")
	}
}

//-----------------Join----------------
func TestGetInnerJoin(t *testing.T) {
	v := new(ToSql)
	n := new(InnerJoin)

	s := v.GetInnerJoin(*n)
	if s != "" {
		t.Errorf("failed to get InnerJoin ")
	}
}
func TestGetOuterJoin(t *testing.T) {
	v := new(ToSql)
	n := new(OuterJoin)

	s := v.GetOuterJoin(*n)
	if s != "" {
		t.Errorf("failed to get OuterJoin ")
	}
}
func TestGetStringJoin(t *testing.T) {
	v := new(ToSql)
	n := new(StringJoin)

	s := v.GetStringJoin(*n)
	if s != "" {
		t.Errorf("failed to get StringJoin ")
	}
}

//-----------------Unary----------------
func TestGetNot(t *testing.T) {
	v := new(ToSql)
	n := new(Not)

	n.expression = Equality{Binary{Literal{1}, Literal{2}}}

	s := v.GetNot(*n)
	if s != "NOT (1 = 2)" {
		t.Log(s)
		t.Errorf("failed to get Not ")
	}
}

func TestGetLock(t *testing.T) {
	v := new(ToSql)
	n := new(Lock)

	s := v.GetLock(*n)
	if s != "" {
		t.Errorf("failed to get Lock ")
	}
}

func TestGetOffset(t *testing.T) {
	v := new(ToSql)
	n := new(Offset)

	n.expression = Literal{10}

	s := v.GetOffset(*n)
	if s != "OFFSET 10" {
		t.Log(s)
		t.Errorf("failed to get Offset ")
	}
}

func TestGetLimit(t *testing.T) {
	v := new(ToSql)
	n := new(Limit)

	n.expression = Literal{100}

	s := v.GetLimit(*n)
	if s != "LIMIT 100" {
		t.Log(s)
		t.Errorf("failed to get Limit ")
	}
}

func TestGetTop(t *testing.T) {
	v := new(ToSql)
	n := new(Top)

	s := v.GetTop(*n)
	if s != "" {
		t.Errorf("failed to get Top ")
	}
}

func TestGetHaving(t *testing.T) {
	v := new(ToSql)
	n := new(Having)

	n.expression = Equality{Binary{Literal{1}, Literal{2}}}

	s := v.GetHaving(*n)
	if s != "HAVING (1 = 2)" {
		t.Log(s)
		t.Errorf("failed to get Having ")
	}
}

func TestGetUnqualifiedColumn(t *testing.T) {
	v := new(ToSql)
	n := new(UnqualifiedColumn)

	s := v.GetUnqualifiedColumn(*n)
	if s != "" {
		t.Errorf("failed to get UnqualifiedColumn ")
	}
}
func TestGetGroup(t *testing.T) {
	v := new(ToSql)
	n := new(Group)

	s := v.GetGroup(*n)
	if s != "" {
		t.Errorf("failed to get Group ")
	}
}

func TestGetGrouping(t *testing.T) {
	v := new(ToSql)
	n := new(Grouping)
	eq1 := Equality{Binary{Literal{1}, Literal{2}}}
	eq2 := Equality{Binary{Literal{3}, Literal{4}}}
	and := And{[]Node{eq1, eq2}}
	n.expression = and
	s := v.GetGrouping(*n)
	if s != "(1 = 2 AND 3 = 4)" {
		t.Log(s)
		t.Errorf("failed to get Grouping ")
	}
}

func TestGetOn(t *testing.T) {
	v := new(ToSql)
	n := new(On)

	n.expression = Equality{Binary{Literal{1}, Literal{2}}}

	s := v.GetOn(*n)
	if s != "ON 1 = 2" {
		t.Log(s)
		t.Errorf("failed to get On ")
	}
}